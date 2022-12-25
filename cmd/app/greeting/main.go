package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"shimabox/example-s3-to-lambda-go/internal/cmd/app/greeting"
)

func main() {
	lambda.Start(greeting.Handler)
}
