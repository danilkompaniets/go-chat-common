package objectStorage

import (
	"context"
	"io"
)

type IS3Manager interface {
	SendFile(ctx context.Context, bucketName, objectName, contentType string, file io.Reader) (string, error)
}
