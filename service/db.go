package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db = mustInitDB("root:NxD5GEARjP@tcp(mysql.spai.svc)/face_recognition")
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
