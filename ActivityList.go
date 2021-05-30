package main

type ActivityList []ActivityDetails

func (al ActivityList) buildStatsFromActivityList(name string, ID string) Stats {
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
