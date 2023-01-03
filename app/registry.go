package app

import (
	"os"

	repositoryInterface "shimabox/example-s3-to-lambda-go/app/domain/repository"
	"shimabox/example-s3-to-lambda-go/app/gateway/storage"
	"shimabox/example-s3-to-lambda-go/app/handler"
	"shimabox/example-s3-to-lambda-go/app/infra/repository"
)

type Registry interface {
	NewGreetingRepository() repositoryInterface.GreetingRepository
	NewGreetingHandler(rep repositoryInterface.GreetingRepository) handler.GreetingHandler
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewGreetingRepository() repositoryInterface.GreetingRepository {
	adapter := storage.NewS3Adapter()
	bucket := storage.NewBucket(os.Getenv("S3_INPUT_BUCKET"))
	storage := storage.NewS3Storage(adapter, bucket)

	return repository.NewGreetingRepository(storage)
}

func (r *registry) NewGreetingHandler(rep repositoryInterface.GreetingRepository) handler.GreetingHandler {
	return handler.NewGreetingHandler(rep)
}
