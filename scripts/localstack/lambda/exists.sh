#!/bin/bash

exists_lambda () {
  local func=$(aws lambda list-functions \
          --endpoint="$ENDPOINT_URL" \
          --query "Functions[?FunctionName=='$1']")

  if [ "$func" != "[]" ]; then
    echo 1
  else
    echo 0
  fi
}