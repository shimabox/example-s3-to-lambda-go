package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"shimabox/example-s3-to-lambda-go/app/adapter/domain/repository"
	"shimabox/example-s3-to-lambda-go/app/adapter/external"
	"shimabox/example-s3-to-lambda-go/app/adapter/usecase"
)

type GreetingHandler interface {
	Handler(ctx context.Context, event events.S3Event)
}

type greetingHandler struct {
	greetingRepository repository.GreetingRepository
}

func NewGreetingHandler(greetingRepository repository.GreetingRepository) *greetingHandler {
	return &greetingHandler{greetingRepository}
}

func (h greetingHandler) Handler(ctx context.Context, event events.S3Event) {
	greetingInput := external.NewGreetingInput(&event)
	greetingUsecase := usecase.NewGreetingUsecase()

	err := greetingUsecase.Handle(ctx, greetingInput, h.greetingRepository)
	if err != nil {
		log.Fatal(err)
	}
}
