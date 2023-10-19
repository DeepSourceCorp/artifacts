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

func (*MockCloudStorageClient) UploadDir(_, _, _ string) error {
	log.Println("not uploading object to upstream since the client is being run in tests")
	return nil
}

func (*MockCloudStorageClient) UploadObject(_, _, _ string) (err error) {
	log.Println("not uploading object to upstream since the client is being run in tests")
	return
}

func (*MockCloudStorageClient) GetObjects(_ string, _ string, _ ...string) error {
	log.Println("mock GetObjects() called, returning successful response")
	return nil
}

func (s *MockCloudStorageClient) NewReader(ctx context.Context, bucket string, path string) (io.ReadCloser, error) {
	obj := s.client.Bucket(bucket).Object(path)
	return obj.NewReader(ctx)
}
