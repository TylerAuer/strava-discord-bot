package main

type Stats struct {
	Name string
	ID   string

	AllCount         int
	AllMovingSeconds int
	//MaxHeartRate

	RunCount         int
	RunMovingSeconds int
	RunMeters        float64
	RunElevationGain float64

	RideCount         int
	RideMovingSeconds int
	RideMeters        float64
	RideElevationGain float64

	WalkOrHikeCount         int
	WalkOrHikeMovingSeconds int
	WalkOrHikeMeters        float64
	WalkOrHikeElevationGain float64
}

func compileStatsFromActivities(name string, ID string, a []ActivityDetails) Stats {
	s := Stats{
		Name: name,
		ID:   ID,
	}

	for _, a := range a {
		s.AllCount++
		s.AllMovingSeconds += a.MovingTime

		if a.Type == "Run" {
			s.RunCount++
			s.RunMovingSeconds = a.MovingTime
			s.RunMeters = a.Distance
			s.RunElevationGain = a.TotalElevationGain
		}

		if a.Type == "Ride" {
			s.RideCount++
			s.RideMovingSeconds = a.MovingTime
			s.RideMeters = a.Distance
			s.RideElevationGain = a.TotalElevationGain
		}

		if a.Type == "Walk" || a.Type == "Hike" {
			s.WalkOrHikeCount++
			s.WalkOrHikeMovingSeconds = a.MovingTime
			s.WalkOrHikeMeters = a.Distance
			s.WalkOrHikeElevationGain = a.TotalElevationGain
		}
	}

	return s
}
