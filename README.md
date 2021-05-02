# Read Strava Data and Post to Discord

Things it can do (by priority)
1. Publish milestones
   1. Passed 100mi, 200mi, 300mi by year
2. Publish charts with updates

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

## Activities to Use During Development

```go
// A bike ride by Fred
handleStravaWebhook(`{
  "aspect_type": "create",
  "event_time": 1618705240,
  "object_id": 5145415643,
  "object_type": "activity",
  "owner_id": 23248014,
  "subscription_id": 188592,
  "updates": {}
}`)

// Tyler's run with Jessica
handleStravaWebhook(`{
  "aspect_type": "create",
  "event_time": 1618702283,
  "object_id": 5145296337,
  "object_type": "activity",
  "owner_id": 20419783,
  "subscription_id": 188592,
  "updates": {}
}`)
```
