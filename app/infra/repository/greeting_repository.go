package repository

import (
	"context"
	"encoding/json"

	"shimabox/example-s3-to-lambda-go/app/domain/model"
	"shimabox/example-s3-to-lambda-go/app/domain/repository"
	"shimabox/example-s3-to-lambda-go/app/gateway/storage"
)

type greetingRepository struct {
	storage storage.S3Storage
}

func NewGreetingRepository(storage storage.S3Storage) repository.GreetingRepository {
	return &greetingRepository{
		storage: storage,
	}
}

func (r *greetingRepository) Read(ctx context.Context, key *model.GreetingEventKey) (*model.Greeting, error) {
	body, err := r.storage.Read(ctx, key.Val())
	if err != nil {
		return nil, err
	}

	var greeting model.Greeting
	if err := json.Unmarshal(body, &greeting); err != nil {
		return nil, err
	}

	return &greeting, nil
}
