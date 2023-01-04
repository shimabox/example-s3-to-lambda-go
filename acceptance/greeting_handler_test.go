package acceptance

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"

	"shimabox/example-s3-to-lambda-go/app/adapter"
	"shimabox/example-s3-to-lambda-go/testhelper"
)

func Test_Handler_挨拶が出力される(t *testing.T) {
	registry := adapter.NewRegistry()
	greetingRepository := registry.NewGreetingRepository()
	greetingHandler := registry.NewGreetingHandler(greetingRepository)

	key1 := "test.json"
	s3Entity := events.S3Entity{
		Object: events.S3Object{
			Key: key1,
		},
	}
	eventRecords := []events.S3EventRecord{
		{
			S3: s3Entity,
		},
	}
	s3Event := events.S3Event{Records: eventRecords}

	got := testhelper.ExtractStdout(t, func() {
		greetingHandler.Handler(context.Background(), s3Event)
	})
	want := "山田太郎: こんにちは"

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Handler is mismatch (-want +got):\n%s", diff)
	}
}
