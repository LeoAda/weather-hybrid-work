package main

import (
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
		for k := range outsideSchedule {
			for j := range forecast {
				forecastDate, _ := time.Parse("2006-01-02T15:04", string(forecast[j].Date))
				if weekDate.Day() == forecastDate.Day() {
					if forecast[j].Hour >= outsideSchedule[k].StartHour && forecast[j].Hour <= outsideSchedule[k].EndHour {
						week[i].ApparentTemperatureMean += forecast[j].ApparentTemperature
						if forecast[j].Hour == outsideSchedule[k].StartHour {
							// If the forecast hour is the start hour, calculate the precipitation total based on the start minute
							week[i].PrecipitationTotal += forecast[j].Precipitation * float64((60-outsideSchedule[k].StartMin)/60)
						} else if forecast[j].Hour == outsideSchedule[k].EndHour {
							// If the forecast hour is the end hour, calculate the precipitation total based on the end minute
							week[i].PrecipitationTotal += forecast[j].Precipitation * float64(outsideSchedule[k].EndMin/60)
						} else {
							// If the forecast hour is within the middle hours, add the full precipitation value
							week[i].PrecipitationTotal += forecast[j].Precipitation
						}
					}

				}

			}
			//Calculate average temperature
			week[i].ApparentTemperatureMean = week[i].ApparentTemperatureMean / float64(outsideSchedule[k].EndHour-outsideSchedule[k].StartHour+1)
		}
	}
	return week
}

func GetNearestMonday(day time.Time) time.Time {
	monday := int(7+(1-float64(day.Weekday()))) % 7
	return day.AddDate(0, 0, int(monday))
}
