package greeting_test

import (
	"context"
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
