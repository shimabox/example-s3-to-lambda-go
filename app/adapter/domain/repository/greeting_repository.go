package repository

import (
	"context"

	"shimabox/example-s3-to-lambda-go/app/adapter/domain/model"
)

type GreetingRepository interface {
	Read(ctx context.Context, key *model.GreetingEventKey) (*model.Greeting, error)
}
