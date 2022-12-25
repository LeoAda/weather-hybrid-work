package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type DayRank struct {
	Day                     string
	ApparentTemperatureRank int
	PrecipitationRank       int
	OverallRank             int
}
type DayForecastOutside struct {
	Day                     string
	ApparentTemperatureMean []float64
	PrecipitationTotal      []float64
}

func (d DayForecastOutside) String() string {
	return fmt.Sprintf("Day: %v\nApparent Temperature Mean: %v Precipitation Total: %v\n", d.Day, d.ApparentTemperatureMean, d.PrecipitationTotal)
}

func (d DayRank) String() string {
	day := strings.TrimSpace(d.Day)
	return fmt.Sprintf("%v, Overall Rank: %v", day, d.OverallRank)
}
func PrintDayRank(days []DayRank) {
	for i, day := range days {
		fmt.Printf("%d. %s\n", i+1, day)
	}
}
func SortDaysByOverallRank(days []DayRank) []DayRank {
	// Use the sort package's Slice function to sort the slice of DayRank structs
	sort.Slice(days, func(i, j int) bool {
		// Return true if the overall rank of the ith element is less than the overall rank of the jth element
		return days[i].OverallRank < days[j].OverallRank
	})
	return days
}

func GenerateRankDays(forecasts [5]DayForecastOutside) []DayRank {
	// Create a slice to store the ranks
	ranks := make([]DayRank, len(forecasts))

	// Iterate through the forecasts
	for i, forecast := range forecasts {
		// Calculate the mean apparent temperature and total precipitation for the day
		meanTemp := mean(forecast.ApparentTemperatureMean)
		totalPrecip := sum(forecast.PrecipitationTotal)

		// Initialize the rank for the day
		rank := DayRank{
			Day:                     forecast.Day,
			ApparentTemperatureRank: 1,
			PrecipitationRank:       len(forecasts),
		}

		// Compare the mean temperature and total precipitation to those of the other days
		// to determine the ranks for each day
		for j, otherForecast := range forecasts {
			if i == j {
				continue
			}
			otherMeanTemp := mean(otherForecast.ApparentTemperatureMean)
			otherTotalPrecip := sum(otherForecast.PrecipitationTotal)

			if meanTemp > otherMeanTemp {
				rank.ApparentTemperatureRank++
			}
			if totalPrecip < otherTotalPrecip {
				rank.PrecipitationRank--
			}
		}

		rank.OverallRank = rank.ApparentTemperatureRank + (rank.PrecipitationRank * 2)

		// Add the rank for the day to the slice
		ranks[i] = rank
	}

	return ranks
}

// Helper function to calculate the mean of a slice of floats
func mean(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

// Helper function to sum a slice of floats
func sum(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total
}

func GenerateWeekForecast(forecast []Forecast, outsideSchedule []OutsideSchedule) [5]DayForecastOutside {
	startDate, _ := time.Parse("2006-01-02T15:04", string(forecast[0].Date))
	var week [5]DayForecastOutside
	for i := range week {
		week[i].Day = startDate.AddDate(0, 0, i).Format("2006-01-02")
		week[i].ApparentTemperatureMean = make([]float64, len(outsideSchedule))
		week[i].PrecipitationTotal = make([]float64, len(outsideSchedule))
	}
	for i := range week {
		weekDate, _ := time.Parse("2006-01-02", string(week[i].Day))
		for k := range outsideSchedule {
			for j := range forecast {
				forecastDate, _ := time.Parse("2006-01-02T15:04", string(forecast[j].Date))
				if weekDate.Day() == forecastDate.Day() {
					if forecast[j].Hour >= outsideSchedule[k].StartHour && forecast[j].Hour <= outsideSchedule[k].EndHour {
						week[i].ApparentTemperatureMean[k] += forecast[j].ApparentTemperature
						if forecast[j].Hour == outsideSchedule[k].StartHour {
							// If the forecast hour is the start hour, calculate the precipitation total based on the start minute
							week[i].PrecipitationTotal[k] += forecast[j].Precipitation * float64((60-outsideSchedule[k].StartMin)/60)
						} else if forecast[j].Hour == outsideSchedule[k].EndHour {
							// If the forecast hour is the end hour, calculate the precipitation total based on the end minute
							week[i].PrecipitationTotal[k] += forecast[j].Precipitation * float64(outsideSchedule[k].EndMin/60)
						} else {
							// If the forecast hour is within the middle hours, add the full precipitation value
							week[i].PrecipitationTotal[k] += forecast[j].Precipitation
						}
					}

				}

			}
			//Calculate average temperature
			week[i].ApparentTemperatureMean[k] = week[i].ApparentTemperatureMean[k] / float64(outsideSchedule[k].EndHour-outsideSchedule[k].StartHour+1)
		}
	}
	return week
}

func GetNearestMonday(day time.Time) time.Time {
	monday := int(7+(1-float64(day.Weekday()))) % 7
	return day.AddDate(0, 0, int(monday))
}
