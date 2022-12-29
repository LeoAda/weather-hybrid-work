package main

import (
	"fmt"
	"time"
)

func LocationMenu(config *Config) error {
	Separator()
	location, err := ChooseLocation()
	Separator()
	if err != nil {
		return fmt.Errorf("Error choosing location: %v", err)
	} else {
		config.Location = location
		return config.Save()
	}
}

func ScheduleMenu(config *Config) error {
	Separator()
	schedule, err := AskDaySchedule()
	Separator()
	if err != nil {
		return fmt.Errorf("Error choosing schedule: %v", err)
	} else {
		config.OutsideSchedule = schedule
		return config.Save()
	}
}

func ForecastMenu(config *Config) error {
	now := time.Now()
	monday := GetNearestMonday(now)
	//choose current week or next week
	Separator()
	fmt.Println("1. Current week")
	fmt.Println("2. Next week")
	var s string
	fmt.Scan(&s)
	switch s {
	case "1":
		monday = monday.AddDate(0, 0, -7)
	case "2":
		//do nothing
	default:
		return fmt.Errorf("Invalid input")
	}
	Separator()
	fmt.Println("Forecast for week of monday ", monday.Format("2006-01-02"))
	friday := monday.AddDate(0, 0, 4)
	forecast := GetWeekForecast(config.Location, monday, friday)

	weekForecast := GenerateWeekForecast(forecast, config.OutsideSchedule)
	Separator()
	rank := SortDaysByOverallRank(GenerateRankDays(weekForecast))
	fmt.Println("Recommended days to work at home : ")
	PrintDayRank(rank)
	Separator()
	return nil
}
func Separator() {
	fmt.Println("------------------------------------------------------------")
}
