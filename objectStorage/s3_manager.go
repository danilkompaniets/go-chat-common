package objectStorage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

type S3Client struct {
	client *minio.Client
}

type ClientConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func NewClient(config ClientConfig) (*S3Client, error) {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})

	if err != nil {
		return nil, err
	}

	return &S3Client{client: minioClient}, nil
}

func (s *S3Client) SendFile(ctx context.Context, bucketName, objectName, contentType string, file io.Reader, size int64) (string, error) {
	info, err := s.client.PutObject(ctx, bucketName, objectName, file, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return info.Key, nil
}
