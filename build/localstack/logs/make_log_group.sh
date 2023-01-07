#!/bin/bash

aws logs create-log-group \
    --endpoint-url="$ENDPOINT_URL" \
    --log-group-name=/aws/lambda/"$LOG_GROUP_NAME"
