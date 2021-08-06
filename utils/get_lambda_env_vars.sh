#!/bin/bash

CRON_FILE=$(pwd)/../secrets/cron
GO_FILE=$(pwd)/../secrets/go

# Stop script if any line fails
set -e

# Don't overwrite old env files
if [ -f $CRON_FILE ]
then
  echo "$CRON_FILE already exists. Rename or delete the old file"
  exit 1
elif [ -f $GO_FILE ]
then
  echo "$GO_FILE already exists. Rename or delete the old file"
  exit 1
fi

# Check that jq is installed
if ! command -v jq >/dev/null 2>&1
then
  echo 'jq command not found. Install with "brew install jq"'
  exit 1
fi

# Pulls the environment variables from the AWS Lambda environment
# and writes them to a local file
write_lambda_env_vars_to_file() {
  FUNCTION_NAME=$1
  DESTINATION_FILE=$2
  
  # Pull env vars from function configuration on AWS
  FUNC_CONFIG=$(aws lambda get-function-configuration --function-name $FUNCTION_NAME)
  ENV_JSON=$(echo $FUNC_CONFIG | jq '.Environment.Variables')
  
  # Map from JSON to "KEY=value" format in Bash list
  LIST_OF_ENV_VARS=$(echo $ENV_JSON | jq "to_entries | map(\"\(.key)=\(.value|tostring)\")")

  # Write list to file with clean up
  for ENV_VAR in $LIST_OF_ENV_VARS
  do
    # Strip out the quotes, commas, and brackets
    ENV_VAR=$(echo $ENV_VAR | tr -d "[]\",")
    
    # Add non-blank lines to file in alphabetical order
    if [ ! -z $ENV_VAR ]
    then
      echo $ENV_VAR
    fi
  done | sort > $DESTINATION_FILE
}

# Login to AWS
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 403525920890.dkr.ecr.us-east-2.amazonaws.com

# Write env vars to files
write_lambda_env_vars_to_file strava-to-discord-cron $CRON_FILE
write_lambda_env_vars_to_file strava-to-discord-go $GO_FILE
