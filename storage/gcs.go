package storage

import (
	"context"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type GoogleCloudStorageClient struct {
	client *storage.Client
}

func NewGoogleCloudStorageClient(ctx context.Context, credentialsJSON []byte) (*GoogleCloudStorageClient, error) {
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		return nil, err
	}

	return &GoogleCloudStorageClient{client}, nil
}

func (s *GoogleCloudStorageClient) UploadDir(bucket, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			err = s.UploadDir(bucket, filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
		} else {
			err = s.UploadObjects(bucket, filepath.Join(path, file.Name()))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *GoogleCloudStorageClient) UploadObjects(bucket string, paths ...string) error {
	var wg sync.WaitGroup
	wg.Add(len(paths))

	for _, path := range paths {
		go func(path string) {
			defer wg.Done()

			file, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("error reading file %q: %v", path, err)
				return
			}

			obj := s.client.Bucket(bucket).Object(path)
			w := obj.NewWriter(context.Background())
			if _, err := w.Write(file); err != nil {
				log.Printf("error uploading file %q: %v", path, err)
				return
			}
			if err := w.Close(); err != nil {
				log.Printf("error closing writer for file %q: %v", path, err)
				return
			}
		}(path)
	}

	wg.Wait()

	return nil
}

func (s *GoogleCloudStorageClient) GetDir(bucket, dir string) error {
	query := &storage.Query{Prefix: dir}
	it := s.client.Bucket(bucket).Objects(context.Background(), query)

	for {
		objAttrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		if objAttrs.Name == dir {
			continue
		}

		err = s.GetObjects(bucket, objAttrs.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GoogleCloudStorageClient) GetObjects(bucket string, paths ...string) error {
	for _, path := range paths {
		obj := s.client.Bucket(bucket).Object(path)
		r, err := obj.NewReader(context.Background())
		if err != nil {
			return err
		}

		defer r.Close()

		data, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filepath.Base(path), data, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
