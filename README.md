# Read Strava Data and Post to Discord

Things it that would be fun to add:
1. Publish milestones like passed 100mi, 200mi, 300mi by year
2. Publish charts with updates


## To Do

- Generalize discordgo into type so I don't have to keep calling dg.close()
- Finish logic for challenge posts
- Combine all cron tasks into single lambda -- after confirming you can set up multiple cron tasks -- that takes an input indicating which to do
  - Jessica's text
  - Weekly update
  - Weekly challenge post
  - Daily lazy checks
## Add user workflow

- Have user visit: https://www.strava.com/oauth/authorize?client_id=63983&redirect_uri=http://localhost&response_type=code&scope=read_all,profile:read_all,activity:read_all
- Approve access
- Copy the code from the url they are dropped at (something like: c1d56dee37aa1914fa8d080355584596ccd5c8f6)
- Finish token exchange with following curl command and add the refresh token into env vars

```
curl -X POST https://www.strava.com/api/v3/oauth/token \
  -d client_id=ReplaceWithClientID \
  -d client_secret=ReplaceWithClientSecret \
  -d code=ReplaceWithCode \
  -d grant_type=authorization_code
```