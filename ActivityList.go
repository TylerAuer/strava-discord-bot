package main

import "fmt"

type ActivityList []ActivityDetails

func (al ActivityList) buildStats(name string, ID string) Stats {
	s := Stats{
		Name:         name,
		ID:           ID,
		MaxHeartRate: 0,
	}

	for _, a := range al {
		s.AllCount++
		s.AllMovingSeconds += a.MovingTime
		s.Heartbeats += int(a.AverageHeartrate * float64(a.MovingTime) / 60.0)
		s.Calories += int(a.Calories)

		if int(a.MaxHeartrate) > s.MaxHeartRate {
			s.MaxHeartRate = int(a.MaxHeartrate)
		}

		if a.Type == "Run" {
			s.RunCount++
			s.RunMovingSeconds += a.MovingTime
			s.RunMeters += a.Distance
			s.RunElevationGain += a.TotalElevationGain
		}

		if a.Type == "Ride" {
			s.RideCount++
			s.RideMovingSeconds += a.MovingTime
			s.RideMeters += a.Distance
			s.RideElevationGain += a.TotalElevationGain
		}

		if a.Type == "Walk" || a.Type == "Hike" {
			s.WalkOrHikeCount++
			s.WalkOrHikeMovingSeconds += a.MovingTime
			s.WalkOrHikeMeters += a.Distance
			s.WalkOrHikeElevationGain += a.TotalElevationGain
		}
	}

	return s
}

func (al ActivityList) buildStatsFor(k Kraftee) Stats {
	s := Stats{
		Name:         k.First,
		ID:           k.StravaId,
		MaxHeartRate: 0,
	}

	for _, a := range al {
		if fmt.Sprint(a.Athlete.ID) == k.StravaId {

			s.AllCount++
			s.AllMovingSeconds += a.MovingTime
			s.Heartbeats += int(a.AverageHeartrate * float64(a.MovingTime) / 60.0)
			s.Calories += int(a.Calories)

			if int(a.MaxHeartrate) > s.MaxHeartRate {
				s.MaxHeartRate = int(a.MaxHeartrate)
			}

			if a.Type == "Run" {
				s.RunCount++
				s.RunMovingSeconds += a.MovingTime
				s.RunMeters += a.Distance
				s.RunElevationGain += a.TotalElevationGain
			}

			if a.Type == "Ride" {
				s.RideCount++
				s.RideMovingSeconds += a.MovingTime
				s.RideMeters += a.Distance
				s.RideElevationGain += a.TotalElevationGain
			}

			if a.Type == "Walk" || a.Type == "Hike" {
				s.WalkOrHikeCount++
				s.WalkOrHikeMovingSeconds += a.MovingTime
				s.WalkOrHikeMeters += a.Distance
				s.WalkOrHikeElevationGain += a.TotalElevationGain
			}
		}
	}

	return s
}

// This is a helper method that's useful for checking activity lists when debugging
func (al ActivityList) printActivitySummaries() {
	for i, a := range al {
		msg := fmt.Sprint(i)                                  // Index in list
		msg += " - " + fmt.Sprint(a.ID) + " - "               // activity id
		msg += fmt.Sprintf("%.1f", metersToMiles(a.Distance)) // distance in miles
		fmt.Println(msg)
	}
}
