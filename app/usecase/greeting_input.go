package usecase

import (
	"shimabox/example-s3-to-lambda-go/app/domain/model"
)

type GreetingInput interface {
	EventToGreetingEventList() *model.GreetingEventList
}
