#!/bin/bash

source "$SCRIPT_DIR"/lambda/exists.sh

if [ $(exists_lambda "$LAMBDA_FUNC_NAME") -eq 1 ]; then
    aws lambda delete-function \
        --endpoint-url="$ENDPOINT_URL" \
        --function-name "$LAMBDA_FUNC_NAME"
    echo "Done lambda delete-function $LAMBDA_FUNC_NAME."
fi