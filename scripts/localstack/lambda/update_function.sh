#!/bin/bash

source "$SCRIPT_DIR"/lambda/exists.sh

if [ $(exists_lambda "$LAMBDA_FUNC_NAME") -eq 1 ]; then
  aws lambda update-function-code \
      --endpoint-url="$ENDPOINT_URL" \
      --function-name "$LAMBDA_FUNC_NAME" \
      --s3-bucket "$S3_LAMBDA_STORAGE_BUCKET" \
      --s3-key "$PACKAGE_NAME".zip \
      --publish

  echo "Done lambda update-function-code $LAMBDA_FUNC_NAME."
fi