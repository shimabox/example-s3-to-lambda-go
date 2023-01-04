package usecase

import (
	"shimabox/example-s3-to-lambda-go/app/adapter/domain/model"
)

type GreetingInput interface {
	EventToGreetingEventList() *model.GreetingEventList
}
