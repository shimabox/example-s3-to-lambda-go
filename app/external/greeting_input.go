package external

import (
	"github.com/aws/aws-lambda-go/events"

	"shimabox/example-s3-to-lambda-go/app/domain/model"
	"shimabox/example-s3-to-lambda-go/app/usecase"
)

type greetingInput struct {
	event *events.S3Event
}

func NewGreetingInput(event *events.S3Event) usecase.GreetingInput {
	return &greetingInput{event}
}

func (i greetingInput) EventToGreetingEventList() *model.GreetingEventList {
	records := i.event.Records

	length := len(records)
	if length == 0 {
		return &model.GreetingEventList{}
	}

	var eventList []*model.GreetingEvent
	for i := range records {
		entity := records[i].S3
		eventList = append(
			eventList,
			model.NewGreetingEvent(model.NewGreetingEventKey(entity.Object.Key)),
		)
	}

	return model.NewGreetingEventList(
		eventList,
	)
}
