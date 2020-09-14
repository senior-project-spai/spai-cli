package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	// Auto load from .env
	_ "github.com/joho/godotenv/autoload"
)

// Image is a model for sql image result
type Image struct {
	ID        string
	Path      string
	Timestamp struct {
		Image           int64
		Age             sql.NullInt64
		Gender          sql.NullInt64
		Emotion         sql.NullInt64
		FaceRecognition sql.NullInt64
	}
}

var db *sql.DB

func init() {
	// Load config from environment variables
	dataSourceName, isDataSourceNamePresent := os.LookupEnv("DATA_SOURCE_NAME")
	if !isDataSourceNamePresent {
		panic(fmt.Errorf("DATA_SOURCE_NAME is not provided in environment"))
	}

	// Initialize database connection
	db = mustInitDB(dataSourceName)
}

func mustInitDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// Setting
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// Ping test
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	return db
}

// RemoveImageFromDB remove image from database with the specific id
func RemoveImageFromDB(imageID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("removeImageFromDB: %w", err)
	}

	// Remove result table
	for _, table := range []string{"age", "gender", "emotion", "face_recognition"} {
		if _, err := tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE image_id = ?", table), imageID); err != nil {
			tx.Rollback()
			return fmt.Errorf("removeImge: delete from %s table: %w", table, err)
		}
	}

	// Remove image table (Main)
	if _, err := tx.ExecContext(ctx, "DELETE FROM image WHERE id = ?", imageID); err != nil {
		tx.Rollback()
		return fmt.Errorf("removeImageFromDB: delete from image table: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("removeImageFromDB: commit: %w", err)
	}

	return nil
}

// ListImages list images in SPAI
func ListImages(limit uint) ([]Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set limit
	query := fmt.Sprintf("SELECT id, path, timestamp, age_timestamp, emotion_timestamp, gender_timestamp, face_recognition_timestamp FROM image ORDER BY timestamp DESC LIMIT %d;", limit)
	if limit == 0 {
		query = "SELECT id, path, timestamp, age_timestamp, emotion_timestamp, gender_timestamp, face_recognition_timestamp FROM image ORDER BY timestamp DESC;"
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ListImage: %w", err)
	}
	defer rows.Close()

	images := make([]Image, 0)
	for rows.Next() {
		var image Image
		if err := rows.Scan(&image.ID, &image.Path, &image.Timestamp.Image, &image.Timestamp.Age, &image.Timestamp.Emotion, &image.Timestamp.Gender, &image.Timestamp.FaceRecognition); err != nil {
			return nil, fmt.Errorf("ListImage: %w", err)
		}
		images = append(images, image)
	}

	return images, nil
}

// GetImage returned an image data in DB
func GetImage(imageID string) (Image, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	image := Image{}
	row := db.QueryRowContext(ctx, "SELECT id, path, timestamp, age_timestamp, emotion_timestamp, gender_timestamp, face_recognition_timestamp FROM image WHERE id = ?;", imageID)
	if err := row.Scan(&image.ID, &image.Path, &image.Timestamp.Image, &image.Timestamp.Age, &image.Timestamp.Emotion, &image.Timestamp.Gender, &image.Timestamp.FaceRecognition); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return Image{}, fmt.Errorf("GetImage: %w", err)
		default:
			panic(err)
		}
	}

	return image, nil
}
