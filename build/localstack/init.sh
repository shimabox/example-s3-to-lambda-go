#!/bin/bash

#############################
# aws configure set
#############################
aws configure set default.region "$AWS_DEFAULT_REGION"
aws configure set aws_access_key_id "$AWS_ACCESS_KEY_ID"
aws configure set aws_secret_access_key "$AWS_SECRET_ACCESS_KEY"

#############################
# S3 initialize.
#############################
bash /docker-entrypoint-initaws.d/s3/make_bucket.sh