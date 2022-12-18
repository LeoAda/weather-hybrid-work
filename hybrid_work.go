package main

import (
	"fmt"
	"time"
)

type DayRank struct {
	Day                     string
	ApparentTemperatureRank int
	PrecipitationRank       int
}
type DayForecastOutside struct {
	Day                     string
	ApparentTemperatureMean float64
	PrecipitationTotal      float64
}

//	Week [5]DayRank

//func SuggestWorkDay([]OutsideSchedule, Forecast){

func GenerateWeekRanks(forecast []Forecast, outsideSchedule []OutsideSchedule) [5]DayForecastOutside {
	startDate, _ := time.Parse("2006-01-02T15:04", string(forecast[0].Date))
	var week [5]DayForecastOutside
	for i := range week {
		week[i].Day = startDate.AddDate(0, 0, i).Format("2006-01-02")
		week[i].ApparentTemperatureMean = 0
		week[i].PrecipitationTotal = 0
	}
	for i := range week {
		weekDate, _ := time.Parse("2006-01-02", string(week[i].Day))
		for j := range forecast {
			forecastDate, _ := time.Parse("2006-01-02T15:04", string(forecast[j].Date))
			if weekDate.Day() == forecastDate.Day() {
				for k := range outsideSchedule {
					if forecast[j].Hour == outsideSchedule[k].StartHour || forecast[j].Hour == outsideSchedule[k].EndHour {
						fmt.Println(forecast[j])
						week[i].ApparentTemperatureMean += forecast[j].ApparentTemperature
						if outsideSchedule[k].StartHour == outsideSchedule[k].EndHour {
							week[i].PrecipitationTotal += forecast[j].Precipitation * float64(60/(outsideSchedule[k].EndMin-outsideSchedule[k].StartMin))
						} else {
							if forecast[j].Hour == outsideSchedule[k].StartHour && outsideSchedule[k].StartMin != 0 {
								week[i].PrecipitationTotal += forecast[j].Precipitation * float64(60-outsideSchedule[k].StartMin)
							} else if forecast[j].Hour == outsideSchedule[k].EndHour && outsideSchedule[k].EndMin != 0 {
								week[i].PrecipitationTotal += forecast[j].Precipitation * float64(outsideSchedule[k].EndMin)
							}
						}

					}
				}

			}
		}
	}
	return week
}

func GetNearestMonday(day time.Time) time.Time {
	monday := int(7+(1-float64(day.Weekday()))) % 7
	return day.AddDate(0, 0, int(monday))
}
