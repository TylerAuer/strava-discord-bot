# MongoDB Schema Planning

This is a planning doc for adding a MongoDB instance to this app to support some features.

## What features will utilize MongoDB

I'd like to be able to compute milestones for Kraftees whenever they log a workout. For example:

- Longest run/bike/walk
- Fastest run of length 0-1, 1-2, 2-3 ...etc
- Most run/bike/walk miles in a week
- Total run/bike/walk miles in a year
- Goal setting and reporting: ex: Tyler met his goal of 10 mi run / week
- Weekly challenges - using a keyword in a workout's title will register it in a weekly challenge (ex: max pushups)

## Organization

- Activity document:
  - `id`: strava activity ID
  - `users`: strava user ID
  - `year`: 2021
  - `month`: 1 to 12
  - `week`: 1 to 52
  - `strava_data`: JSON object containing all the data from Strava