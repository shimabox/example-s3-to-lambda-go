#!/bin/bash

bash "$SCRIPT_DIR"/helper/upload.sh

sleep 10 # 反映を待つ

bash "$SCRIPT_DIR"/helper/get_log.sh
