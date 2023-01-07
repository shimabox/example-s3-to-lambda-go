#!/bin/bash

bash "$SCRIPT_DIR"/helper/upload.sh

# TODO: logStreamNameが取れるまでループで確認する
sleep 15 # 反映を待つ

bash "$SCRIPT_DIR"/helper/get_log.sh
