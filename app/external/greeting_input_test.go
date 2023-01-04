package external_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"

	"shimabox/example-s3-to-lambda-go/app/adapter/domain/model"
	"shimabox/example-s3-to-lambda-go/app/external"
)

func Test_EventToGreetingEventList_イベントが存在しないとき初期値のイベントリストモデルが返る(t *testing.T) {
	input := external.NewGreetingInput(&events.S3Event{
		Records: []events.S3EventRecord{},
	})

	want := &model.GreetingEventList{}
	got := input.EventToGreetingEventList()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("EventToGreetingEventList is mismatch (-want +got):\n%s", diff)
	}
}

func Test_EventToGreetingEventList_イベントが存在すればイベントリストモデルを返す(t *testing.T) {
	key1 := "test.json"
	key2 := "test2.json"
	key3 := "test3.json"
	s3Entity := events.S3Entity{
		Object: events.S3Object{
			Key: key1,
		},
	}
	s3Entity2 := events.S3Entity{
		Object: events.S3Object{
			Key: key2,
		},
	}
	s3Entity3 := events.S3Entity{
		Object: events.S3Object{
			Key: key3,
		},
	}
	eventRecords := []events.S3EventRecord{
		{
			S3: s3Entity,
		},
		{
			S3: s3Entity3,
		},
		{
			S3: s3Entity2,
		},
	}

	list := []*model.GreetingEvent{
		model.NewGreetingEvent(model.NewGreetingEventKey(key1)),
		model.NewGreetingEvent(model.NewGreetingEventKey(key3)),
		model.NewGreetingEvent(model.NewGreetingEventKey(key2)),
	}
	want := model.NewGreetingEventList(list)

	input := external.NewGreetingInput(&events.S3Event{
		Records: eventRecords,
	})
	got := input.EventToGreetingEventList()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("EventToGreetingEventList is mismatch (-want +got):\n%s", diff)
	}
}
