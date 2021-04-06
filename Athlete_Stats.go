package main

import "fmt"

type AthleteStats struct {
	BiggestRideDistance       float64 `json:"biggest_ride_distance"`
	BiggestClimbElevationGain float64 `json:"biggest_climb_elevation_gain"`
	RecentRideTotals          struct {
		Count            int     `json:"count"`
		Distance         float64 `json:"distance"`
		MovingTime       int     `json:"moving_time"`
		ElapsedTime      int     `json:"elapsed_time"`
		ElevationGain    float64 `json:"elevation_gain"`
		AchievementCount int     `json:"achievement_count"`
	} `json:"recent_ride_totals"`
	AllRideTotals struct {
		Count         int `json:"count"`
		Distance      int `json:"distance"`
		MovingTime    int `json:"moving_time"`
		ElapsedTime   int `json:"elapsed_time"`
		ElevationGain int `json:"elevation_gain"`
	} `json:"all_ride_totals"`
	RecentRunTotals struct {
		Count            int     `json:"count"`
		Distance         float64 `json:"distance"`
		MovingTime       int     `json:"moving_time"`
		ElapsedTime      int     `json:"elapsed_time"`
		ElevationGain    float64 `json:"elevation_gain"`
		AchievementCount int     `json:"achievement_count"`
	} `json:"recent_run_totals"`
	AllRunTotals struct {
		Count         int `json:"count"`
		Distance      int `json:"distance"`
		MovingTime    int `json:"moving_time"`
		ElapsedTime   int `json:"elapsed_time"`
		ElevationGain int `json:"elevation_gain"`
	} `json:"all_run_totals"`
	RecentSwimTotals struct {
		Count            int     `json:"count"`
		Distance         float64 `json:"distance"`
		MovingTime       int     `json:"moving_time"`
		ElapsedTime      int     `json:"elapsed_time"`
		ElevationGain    float64 `json:"elevation_gain"`
		AchievementCount int     `json:"achievement_count"`
	} `json:"recent_swim_totals"`
	AllSwimTotals struct {
		Count         int `json:"count"`
		Distance      int `json:"distance"`
		MovingTime    int `json:"moving_time"`
		ElapsedTime   int `json:"elapsed_time"`
		ElevationGain int `json:"elevation_gain"`
	} `json:"all_swim_totals"`
	YtdRideTotals struct {
		Count         int `json:"count"`
		Distance      int `json:"distance"`
		MovingTime    int `json:"moving_time"`
		ElapsedTime   int `json:"elapsed_time"`
		ElevationGain int `json:"elevation_gain"`
	} `json:"ytd_ride_totals"`
	YtdRunTotals struct {
		Count         int `json:"count"`
		Distance      int `json:"distance"`
		MovingTime    int `json:"moving_time"`
		ElapsedTime   int `json:"elapsed_time"`
		ElevationGain int `json:"elevation_gain"`
	} `json:"ytd_run_totals"`
	YtdSwimTotals struct {
		Count         int `json:"count"`
		Distance      int `json:"distance"`
		MovingTime    int `json:"moving_time"`
		ElapsedTime   int `json:"elapsed_time"`
		ElevationGain int `json:"elevation_gain"`
	} `json:"ytd_swim_totals"`
}

func (as AthleteStats) YtdRunsTotalsString() string {
	return fmt.Sprint(as.YtdRunTotals.Count) + " runs over " + fmt.Sprintf("%f.1", metersToMiles(as.YtdRideTotals.Distance)) + " in " + fmt.Sprintf("%f.1", secondsToHours(as.YtdRideTotals.ElapsedTime))
}
