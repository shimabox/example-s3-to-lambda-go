package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"shimabox/example-s3-to-lambda-go/app/domain/model"
)

func Test_Exists_イベントが存在しなければfalse(t *testing.T) {
	var events []*model.GreetingEvent
	eventList := model.NewGreetingEventList(events)

	want := false
	got := eventList.Exists()

	if equal := cmp.Equal(want, got); !equal {
		t.Errorf("Exists() = %v, want %v", got, want)
	}
}

func Test_Exists_イベントが存在すればtrue(t *testing.T) {
	events := []*model.GreetingEvent{
		{
			EventKey: model.NewGreetingEventKey("test"),
		},
	}
	eventList := model.NewGreetingEventList(events)

	want := true
	got := eventList.Exists()

	if equal := cmp.Equal(want, got); !equal {
		t.Errorf("Exists() = %v, want %v", got, want)
	}
}

func Test_GetFirstEvent_イベントが取得できる(t *testing.T) {
	event := model.NewGreetingEvent(model.NewGreetingEventKey("test"))
	events := []*model.GreetingEvent{
		event,
	}
	eventList := model.NewGreetingEventList(events)

	want := event
	got := eventList.GetFirstEvent()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetFirstEvent is mismatch (-want +got):\n%s", diff)
	}
}

func Test_GetFirstEvent_複数イベントがある場合最初のイベントだけ取得できる(t *testing.T) {
	event := model.NewGreetingEvent(model.NewGreetingEventKey("test"))
	event2 := model.NewGreetingEvent(model.NewGreetingEventKey("test2"))
	events := []*model.GreetingEvent{
		event,
		event2,
	}
	eventList := model.NewGreetingEventList(events)

	want := event
	got := eventList.GetFirstEvent()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetFirstEvent is mismatch (-want +got):\n%s", diff)
	}
}
