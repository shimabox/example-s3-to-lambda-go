package greeting_test

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"

	"shimabox/example-s3-to-lambda-go/internal/cmd/app/greeting"
)

func Test_Handler(t *testing.T) {
	type args struct {
		in0   context.Context
		event events.S3Event
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			greeting.Handler(tt.args.in0, tt.args.event)
		})
	}
}
