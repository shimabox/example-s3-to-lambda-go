#!/bin/bash

#############################
# S3 initialize.
#############################
bash /docker-entrypoint-initaws.d/s3/make_bucket.sh

#############################
# CloudWatch initialize.
#############################
bash /docker-entrypoint-initaws.d/logs/make_log_group.sh
