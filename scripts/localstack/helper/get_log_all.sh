#!/bin/bash

LOG_GROUP=/aws/lambda/example-s3-to-lambda-go
LOG_STREAMS=$(aws logs describe-log-streams \
  --endpoint-url="$ENDPOINT_URL" \
  --log-group-name "$LOG_GROUP" \
  --query logStreams[].logStreamName \
  --output text)

list=($(echo $LOG_STREAMS)) # 半角スペースで配列につめる
for i in $(seq 1 ${#list[@]})
do
  echo "----- $i -----"
  aws logs get-log-events \
      --endpoint-url="$ENDPOINT_URL" \
      --log-group-name "$LOG_GROUP" \
      --log-stream-name "${list[$i-1]}" \
      --query events[].message \
      --output json
done