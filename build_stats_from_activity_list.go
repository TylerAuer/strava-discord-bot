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
