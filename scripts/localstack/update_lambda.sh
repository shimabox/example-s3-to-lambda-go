#!/bin/bash

#############################
# S3.
#############################
bash "$SCRIPT_DIR"/s3/upload_lambda.sh

#############################
# Lambda
#############################
bash "$SCRIPT_DIR"/lambda/update_function.sh