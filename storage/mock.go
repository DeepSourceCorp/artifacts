package storage

import (
	"context"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type MockCloudStorageClient struct {
	client *storage.Client
}

func NewMockCloudStorageClient(ctx context.Context, credentialsJSON []byte) (*MockCloudStorageClient, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		return nil, err
	}
	return &MockCloudStorageClient{client}, nil
}

func (s *MockCloudStorageClient) UploadDir(bucket, src, dst string) error {
	log.Println("not uploading object to upstream since the client is being run in tests")
	return nil
}

func (s *MockCloudStorageClient) UploadObject(bucket, src, dst string) (err error) {
	log.Println("not uploading object to upstream since the client is being run in tests")
	return
}

func (s *MockCloudStorageClient) GetObjects(bucket string, destinationPath string, paths ...string) error {
	log.Println("mock GetObjects() called, returning successful response")
	return nil
}

func (s *MockCloudStorageClient) NewReader(ctx context.Context, bucket string, path string) (io.ReadCloser, error) {
	obj := s.client.Bucket(bucket).Object(path)
	return obj.NewReader(ctx)
}
