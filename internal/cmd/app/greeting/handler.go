package greeting

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Handler(_ context.Context, event events.S3Event) {
	if len(event.Records) == 0 {
		return
	}

	ses, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	config := getConfig()
	s3Ses := s3.New(ses, &config)

	record := event.Records[0]
	s3Entity := record.S3
	bucket := s3Entity.Bucket.Name
	key := s3Entity.Object.Key

	obj, err := s3Ses.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatal(err)
	}

	body, err := read(obj)
	if err != nil {
		log.Fatal(err)
	}

	var myEvent MyEvent
	if err := json.Unmarshal(body, &myEvent); err != nil {
		log.Fatal(err)
	}

	myEvent.SayHello()
}

func getConfig() aws.Config {
	c := aws.Config{}
	if os.Getenv("IS_LOCAL") == "true" {
		c.Endpoint = aws.String("http://localstack:4566")
		c.S3ForcePathStyle = aws.Bool(true)
	}

	return c
}

func read(o *s3.GetObjectOutput) ([]byte, error) {
	defer o.Body.Close()

	b, err := io.ReadAll(o.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
