#!/bin/bash

aws s3 cp "$SCRIPT_DIR"/helper/test.json s3://"$S3_INPUT_BUCKET" \
    --endpoint-url="$ENDPOINT_URL"
ret_code=${?}

if [ "$ret_code" -eq 0 ]; then
  sleep 3 # ログへの反映を待つ
  echo "Put to S3 done."
fi
