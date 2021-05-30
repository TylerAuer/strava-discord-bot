package main

type Stats struct {
	Name string
	ID   string

	AllCount         int
	AllMovingSeconds int
	Heartbeats       int
	MaxHeartRate     int
	Calories         int

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

func buildStatsFromActivityList(name string, ID string, a []ActivityDetails) Stats {
	s := Stats{
		Name:         name,
		ID:           ID,
		MaxHeartRate: 0,
	}

	for _, a := range a {
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
