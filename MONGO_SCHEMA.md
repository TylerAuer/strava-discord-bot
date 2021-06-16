# MongoDB Schema Planning

This is a planning doc for adding a MongoDB instance to this app to support some features.

## Why is state needed?

Currently `strava-discord-bot` is stateless. It repulls all the data it needs everytime it is triggered by the Strava webhook. Given the limited number of users (10) and the limited number of webhook triggers per day (usually < 20) there's no risk of exceeding the Strava API limits. However, I'd like to add some new features that would require more API calls. It would probably still be within Strava's limits, but this is a good opportunity to play with a new (to me) technology.

## What features will utilize MongoDB

I'd like to be able to compute milestones for Kraftees whenever they log a workout. For example:

- Longest run/bike/walk
- Fastest run of length 0-1, 1-2, 2-3 ...etc
- Most run/bike/walk miles in a week
- Total run/bike/walk miles in a year
- Goal setting and reporting: ex: Tyler met his goal of 10 mi run / week
- Weekly challenges - using a keyword in a workout's title will register it in a weekly challenge (ex: max pushups)

## Schema

