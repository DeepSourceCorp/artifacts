package storage

import (
	"context"
	"testing"

	"github.com/minio/minio-go/v7"
)

var testS3Client *S3StorageClient

func TestMain(m *testing.M) {
	initS3Client()
	m.Run()
	cleanup()
}

func initS3Client() {
	if testS3Client == nil {
		testS3Client, _ = NewS3StorageClient(context.Background(), "localhost:9000", "minioadmin", "minioadmin", false)
	}

	testS3Client.minioClient.MakeBucket(context.Background(), "testbucket", minio.MakeBucketOptions{})
}

func cleanup() {
	testS3Client.minioClient.RemoveBucketWithOptions(context.Background(), "testbucket", minio.RemoveBucketOptions{ForceDelete: true})
	testS3Client = nil
}

func TestS3StorageClient_UploadObject(t *testing.T) {
	type fields struct {
		minioClient *minio.Client
	}
	type args struct {
		bucket string
		src    string
		dst    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"No error",
			fields{testS3Client.minioClient},
			args{"testbucket", "storage.go", "testdata"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &S3StorageClient{
				minioClient: tt.fields.minioClient,
			}
			if err := s.UploadObject(tt.args.bucket, tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("S3StorageClient.UploadObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


