package greeting_test

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"

	"shimabox/example-s3-to-lambda-go/internal/cmd/greeting"
	"shimabox/example-s3-to-lambda-go/testhelper"
)

func Test_Handler_イベント対象のレコードが存在しない場合_何も出力されない(t *testing.T) {
	f := func() {
		greeting.Handler(context.Background(), events.S3Event{})
	}

	got := testhelper.ExtractStdout(t, f)
	want := ""

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Handler is mismatch (-want +got):\n%s", diff)
	}
}

func Test_Handler_イベント対象のレコードが存在する場合_レコードの中身が出力される(t *testing.T) {
	f := func() {
		// NOTE: S3_INPUT_BUCKETにtest.jsonがuploadされていることがテストを通す条件
		// TODO: 前提条件が邪魔なのでいずれうまいことやる
		s3Entity := events.S3Entity{
			Bucket: events.S3Bucket{
				Name: os.Getenv("S3_INPUT_BUCKET"),
			},
			Object: events.S3Object{
				Key: "test.json",
			},
		}
		eventRecords := []events.S3EventRecord{
			{
				S3: s3Entity,
			},
		}
		greeting.Handler(context.Background(), events.S3Event{
			Records: eventRecords,
		})
	}

	got := testhelper.ExtractStdout(t, f)
	want := "Hello 山田太郎: こんにちは"

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Handler is mismatch (-want +got):\n%s", diff)
	}
}
