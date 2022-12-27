#!/bin/bash

#############################
# S3.
#############################
bash "$SCRIPT_DIR"/s3/upload_lambda.sh
bash "$SCRIPT_DIR"/s3/add_notification_event.sh

#############################
# Lambda
#############################
bash "$SCRIPT_DIR"/lambda/delete_function.sh
bash "$SCRIPT_DIR"/lambda/create_function.sh
bash "$SCRIPT_DIR"/lambda/update_function_configuration.sh