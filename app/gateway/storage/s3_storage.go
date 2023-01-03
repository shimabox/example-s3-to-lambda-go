package storage

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type bucket struct {
	name string
}

func NewBucket(name string) *bucket {
	return &bucket{name}
}

type S3Adapter interface {
	Open() (*s3.S3, error)
}

type s3Adapter struct{}

func NewS3Adapter() S3Adapter {
	return &s3Adapter{}
}

type S3Storage interface {
	Read(_ context.Context, key string) ([]byte, error)
	ReadBody(o *s3.GetObjectOutput) ([]byte, error)
}

type s3Storage struct {
	adapter S3Adapter
	bucket  *bucket
}

func NewS3Storage(adapter S3Adapter, bucket *bucket) S3Storage {
	return &s3Storage{
		adapter,
		bucket,
	}
}

func (s *s3Storage) Read(_ context.Context, key string) ([]byte, error) {
	client, err := s.adapter.Open()
	if err != nil {
		return nil, err
	}

	obj, err := client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket.name),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	body, err := s.ReadBody(obj)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *s3Storage) ReadBody(o *s3.GetObjectOutput) ([]byte, error) {
	defer o.Body.Close()

	b, err := io.ReadAll(o.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *s3Adapter) Open() (*s3.S3, error) {
	ses, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	config := aws.Config{}
	if os.Getenv("IS_LOCAL") == "true" {
		config.Endpoint = aws.String(os.Getenv("LOCALSTACK_URL"))
		config.S3ForcePathStyle = aws.Bool(true)
		config.Region = aws.String(os.Getenv("AWS_DEFAULT_REGION"))
	}

	return s3.New(ses, &config), nil
}
