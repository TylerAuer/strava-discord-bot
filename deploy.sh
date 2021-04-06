docker build -t strava-discord-bot . 
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 403525920890.dkr.ecr.us-east-2.amazonaws.com
docker tag  strava-discord-bot:latest 403525920890.dkr.ecr.us-east-2.amazonaws.com/strava-discord-bot:latest
docker push 403525920890.dkr.ecr.us-east-2.amazonaws.com/strava-discord-bot:latest