package greeting_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"shimabox/example-s3-to-lambda-go/internal/cmd/greeting"
	"shimabox/example-s3-to-lambda-go/testhelper"
)

func Test_SayHello(t *testing.T) {
	myEvent := greeting.MyEvent{
		Name:    "テスト太郎",
		Message: "こんにちは",
	}
	got := testhelper.ExtractStdout(t, myEvent.SayHello)
	want := "Hello テスト太郎: こんにちは"

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("SayHello is mismatch (-want +got):\n%s", diff)
	}
}
