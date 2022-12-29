#!/bin/bash

aws s3 cp "$DIST_DIR"/"$PACKAGE_NAME".zip s3://"$S3_LAMBDA_STORAGE_BUCKET" \
    --endpoint-url="$ENDPOINT_URL"