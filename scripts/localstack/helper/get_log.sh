#!/bin/bash

LOG_GROUP=/aws/lambda/example-s3-to-lambda-go
LOG_STREAM=$(aws logs describe-log-streams \
  --endpoint-url="$ENDPOINT_URL" \
  --log-group-name "$LOG_GROUP" \
  --max-items 1 \
  --order-by LastEventTime \
  --descending \
  --query logStreams[].logStreamName \
  --output text | head -n 1)

aws logs get-log-events \
    --endpoint-url="$ENDPOINT_URL" \
    --log-group-name "$LOG_GROUP" \
    --log-stream-name "$LOG_STREAM" \
    --query events[].message \
    --output json
