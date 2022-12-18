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

	//schedule := AskDaySchedule()
	//config.OutsideSchedule = schedule
	//SaveConfig(config)
	fmt.Println(config.OutsideSchedule)

	// Get forecast
	now := time.Now()
	monday := GetNearestMonday(now)
	forecast := GetWeekForecast(config.Location, monday)
	fmt.Println(forecast)

	fmt.Println(GenerateWeekRanks(forecast, config.OutsideSchedule))

}
