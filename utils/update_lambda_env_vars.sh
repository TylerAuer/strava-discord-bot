#!/bin/bash

CRON_FILE=$(pwd)/../secrets/cron
GO_FILE=$(pwd)/../secrets/go

# Stop script if any line fails
set -e

# Be sure that the local env files exist
if [ ! -f $CRON_FILE ]
then
  echo "$CRON_FILE does not exists. Try running get_lambda_env_vars to create it."
  exit 1
elif [ ! -f $GO_FILE ]
then
  echo "$GO_FILE does not exists. Try running get_lambda_env_vars to create it."
  exit 1
fi

upload_env_vars_to_aws_lambda() {
  FUNCTION_NAME=$1
  FILE=$2
  
  COMMA_SEPARATED_ENV_VARS_AS_STRING=$(paste -sd "," $FILE)

  # Get the env vars
  aws lambda update-function-configuration --function-name $FUNCTION_NAME \
      --environment "Variables={$COMMA_SEPARATED_ENV_VARS_AS_STRING}"
}

# Login to AWS
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 403525920890.dkr.ecr.us-east-2.amazonaws.com

# Update Lambda environment variables
upload_env_vars_to_aws_lambda "strava-to-discord-go" $GO_FILE
upload_env_vars_to_aws_lambda "strava-to-discord-cron" $CRON_FILE