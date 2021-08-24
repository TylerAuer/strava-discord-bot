#!/bin/bash

# Login to AWS
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 403525920890.dkr.ecr.us-east-2.amazonaws.com

aws logs get-log-events \
  --log-group-name /aws/lambda/strava-to-discord-go \
  --log-stream-name $(cat out) \
  --limit 5
