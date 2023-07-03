package storage

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
)

type MockCloudStorageClient struct {
	client *storage.Client
}

func NewMockCloudStorageClient(ctx context.Context, _ []byte) (*MockCloudStorageClient, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &MockCloudStorageClient{client}, nil
}

func (s *MockCloudStorageClient) UploadDir(bucket, src, dst string) error {
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			err = s.UploadDir(bucket, filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			if err != nil {
				log.Println(err)
				return err
			}
		} else {
			err = s.UploadObject(bucket, filepath.Join(src, file.Name()), filepath.Join(dst, file.Name()))
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	return nil
}

func (s *MockCloudStorageClient) UploadObject(bucket, src, dst string) (err error) {
	log.Println("not uploading object to upstream since the client is being run in tests")
	return
}

func (s *MockCloudStorageClient) GetObjects(bucket string, destinationPath string, paths ...string) error {
	for _, path := range paths {
		obj := s.client.Bucket(bucket).Object(path)
		r, err := obj.NewReader(context.Background())
		if err != nil {
			return err
		}

		defer r.Close()

		data, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		err = os.WriteFile(destinationPath, data, 0o644)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *MockCloudStorageClient) NewReader(ctx context.Context, bucket string, path string) (io.ReadCloser, error) {
	obj := s.client.Bucket(bucket).Object(path)
	return obj.NewReader(ctx)
}
