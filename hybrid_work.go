package main

import (
	"time"
)

type DayRank struct {
	Day                     string
	ApparentTemperatureRank int
	PrecipitationRank       int
}

//	Week []DayRank

//func SuggestWorkDay([]OutsideSchedule, Forecast){
/*
func GenerateWeekRanks(forecast Forecast, outsideSchedule []OutsideSchedule) []DayRank {
	start_date, _ := time.Parse("2022-12-19T02:00", string(forecast.Time[0]))
	for day := 0; day < 7; day++ {
		for index, schedule := range outsideSchedule {
			start_time = time.Time.Day(start_date.Year(), start_date.Month(), start_date.Day(), schedule.StartTime.Hour(), schedule.StartTime.Minute(), 0, 0, time.UTC)

		}
	}
}
*/
func GetNearestMonday(day time.Time) time.Time {
	monday := int(7+(1-float64(day.Weekday()))) % 7
	return day.AddDate(0, 0, int(monday))
}
