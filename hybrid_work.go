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
		for k := range outsideSchedule {
			for j := range forecast {
				forecastDate, _ := time.Parse("2006-01-02T15:04", string(forecast[j].Date))
				if weekDate.Day() == forecastDate.Day() {
					fmt.Println(forecast[j])
					if forecast[j].Hour == outsideSchedule[k].StartHour && forecast[j].Hour == outsideSchedule[k].EndHour {
						//same hour
						week[i].ApparentTemperatureMean += forecast[j].ApparentTemperature
						week[i].PrecipitationTotal += forecast[j].Precipitation * float64(60/(outsideSchedule[k].EndMin-outsideSchedule[k].StartMin))
					} else if forecast[j].Hour == outsideSchedule[k].StartHour && forecast[j].Hour != outsideSchedule[k].EndHour {
						//start hour
						week[i].ApparentTemperatureMean += forecast[j].ApparentTemperature
						week[i].PrecipitationTotal += forecast[j].Precipitation * float64(outsideSchedule[k].StartMin)
					} else if forecast[j].Hour != outsideSchedule[k].StartHour && forecast[j].Hour == outsideSchedule[k].EndHour {
						//end hour
						week[i].ApparentTemperatureMean += forecast[j].ApparentTemperature
						week[i].PrecipitationTotal += forecast[j].Precipitation * float64(outsideSchedule[k].EndMin)
					} else if forecast[j].Hour > outsideSchedule[k].StartHour && forecast[j].Hour < outsideSchedule[k].EndHour {
						//middle hours
						week[i].ApparentTemperatureMean += forecast[j].ApparentTemperature
						week[i].PrecipitationTotal += forecast[j].Precipitation
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
