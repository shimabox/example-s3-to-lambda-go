#!/bin/bash

aws s3api put-bucket-notification-configuration \
    --bucket "$S3_INPUT_BUCKET" \
    --endpoint-url="$ENDPOINT_URL" \
    --notification-configuration "{
      \"LambdaFunctionConfigurations\": [
        {
          \"LambdaFunctionArn\": \"arn:aws:lambda:ap-northeast-1:000000000000:function:$LAMBDA_FUNC_NAME\",
          \"Events\": [\"s3:ObjectCreated:Put\"]
        }
    }"