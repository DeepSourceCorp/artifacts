package storage

import (
	"context"
	"fmt"
	"io"
)

type StorageClient interface {
	UploadDir(string, string, string) error
	UploadObjects(string, ...string) error
	UploadObject(string, string, string) error
	GetDir(string, string) error
	GetObjects(string, string, ...string) error
	NewReader(context.Context, string, string) (io.ReadCloser, error)
}

func NewStorageClient(ctx context.Context, storageType string, credentials []byte) (StorageClient, error) {
	if storageType == "gcs" {
		return NewGoogleCloudStorageClient(ctx, credentials)
	}

	return &GoogleCloudStorageClient{}, fmt.Errorf("expected storageType to be 'gcs'. Received %s", storageType)
}
