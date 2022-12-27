#!/bin/bash

aws s3 mb s3://"$S3_LAMBDA_STORAGE_BUCKET" \
    --endpoint-url="$ENDPOINT_URL"
aws s3 mb s3://"$S3_INPUT_BUCKET" \
    --endpoint-url="$ENDPOINT_URL"