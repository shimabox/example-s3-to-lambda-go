#!/bin/bash

aws s3 cp "$SCRIPT_DIR"/helper/test.json s3://"$S3_INPUT_BUCKET" \
    --endpoint-url="$ENDPOINT_URL"
