package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	/*
		config := Config{}
		SaveConfig(config)
		location := ChooseLocation()
		fmt.Println(location)
		config.Location = location
		SaveConfig(config)
	*/
	config, _ := LoadConfig()
	fmt.Println(config.Location)

	schedule := AskDaySchedule()
	config.OutsideSchedule = schedule
	SaveConfig(config)
	PrintSchedule(config.OutsideSchedule)

	// Get forecast
	now := time.Now()
	fmt.Println(now)
	fmt.Println(GetNearestMonday(now))
	monday := GetNearestMonday(now)
	forecast := GetWeekForecast(config.Location, monday)
	fmt.Println(forecast)

}
