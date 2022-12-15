package main

import (
	"fmt"
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

	// Get forecast
	forecast := GetForecast(config.Location)
	fmt.Println(forecast)

	//schedule := AskDaySchedule()
	//config.OutsideSchedule = schedule
	//SaveConfig(config)
	PrintSchedule(config.OutsideSchedule)

}
