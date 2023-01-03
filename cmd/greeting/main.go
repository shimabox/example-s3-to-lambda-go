package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"shimabox/example-s3-to-lambda-go/app"
)

func main() {
	registry := app.NewRegistry()
	greetingRepository := registry.NewGreetingRepository()
	greetingHandler := registry.NewGreetingHandler(greetingRepository)

	lambda.Start(greetingHandler.Handler)
}
