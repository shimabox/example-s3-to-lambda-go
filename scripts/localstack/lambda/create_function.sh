#!/bin/bash

source "$SCRIPT_DIR"/lambda/exists.sh

if [ $(exists_lambda "$LAMBDA_FUNC_NAME") -eq 0 ]; then
    aws lambda create-function \
        --endpoint-url="$ENDPOINT_URL" \
        --function-name "$LAMBDA_FUNC_NAME" \
        --role arn:aws:iam::000000000000:role/dummy-lambda-role \
        --handler greeting_handler \
        --runtime go1.x \
        --code S3Bucket="$S3_LAMBDA_STORAGE_BUCKET",S3Key=greeting_handler.zip \
        --timeout 15 \
        --memory-size 512 \
        --package-type Zip
    echo "Done lambda create-function $LAMBDA_FUNC_NAME."
fi