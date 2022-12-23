package main

import (
	"fmt"
	"time"
)

// https://github.com/Pungyeon/clean-go-article
func main() {
	fmt.Println("Welcome to the weather hybrid work app")
	config, _ := LoadConfig()

	//Menu(config)
	fmt.Println(config.Location)

	schedule := AskDaySchedule()
	config.OutsideSchedule = schedule
	config.Save()
	fmt.Println(config.OutsideSchedule)

	// Get forecast
	now := time.Now()
	monday := GetNearestMonday(now)
	fmt.Println("Forecast for week of monday ", monday.Format("2006-01-02"))
	forecast := GetWeekForecast(config.Location, monday)
	//fmt.Println(forecast)

	weekForecast := GenerateWeekForecast(forecast, config.OutsideSchedule)
	fmt.Println(weekForecast)

	rank := GenerateRankDays(weekForecast)
	fmt.Println(rank)

	rank = SortDaysByOverallRank(rank)
	fmt.Println(rank)

}

func Menu(config Config) {
	var s string
	var err error
	fmt.Println("1. Choose location")
	fmt.Println("2. Choose schedule")
	fmt.Println("3. Get forecast")
	fmt.Println("4. Exit")
	fmt.Scan(&s)
	switch s {
	case "1":
		err = LocationMenu(config)
	case "2":
		//AskDaySchedule()
	case "3":
		//GetForecast()
	case "4":
		//Exit()
	default:
		fmt.Println("Invalid input")
	}
	if err != nil {
		fmt.Println(err)
	}
}
