#!/bin/bash

aws s3api put-bucket-notification-configuration \
    --bucket "$S3_INPUT_BUCKET" \
    --endpoint-url="$ENDPOINT_URL" \
    --notification-configuration file://"$SCRIPT_DIR"/s3/s3_notification_config.json