#!/bin/bash

source "$SCRIPT_DIR"/lambda/exists.sh

if [ $(exists_lambda "$LAMBDA_FUNC_NAME") -eq 1 ]; then
    aws lambda update-function-configuration \
        --endpoint-url="$ENDPOINT_URL" \
        --function-name "$LAMBDA_FUNC_NAME" \
        --environment Variables={IS_LOCAL=true}

    echo "Done lambda update-function-configuration $LAMBDA_FUNC_NAME."
fi