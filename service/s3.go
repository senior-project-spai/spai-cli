package service

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var s3 *minio.Client

func init() {
	s3 = mustInitS3()
}

func mustInitS3() *minio.Client {
	// WARNING: not throw error if endpoint is wrong
	s3, err := minio.New("minio.spai.svc:9000", &minio.Options{
		Creds: credentials.NewStaticV4("minio", "minio123", ""),
	})
	if err != nil {
		panic(err)
	}

	return s3
}

// RemoveImageFromS3 remove image from S3 storage with a specific id
func RemoveImageFromS3(imageID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Remove image from S3
	if err := s3.RemoveObject(ctx, "image", imageID, minio.RemoveObjectOptions{GovernanceBypass: true}); err != nil {
		return fmt.Errorf("removeImageFromS3: %w", err)
	}

	return nil
}
