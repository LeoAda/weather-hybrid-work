package main

import (
	"fmt"
)

// https://github.com/Pungyeon/clean-go-article
func main() {
	fmt.Println("Welcome to the weather hybrid work app")
	config, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	Menu(&config)

	fmt.Println(config.Location)
	fmt.Println(config.OutsideSchedule)

}

func Menu(config *Config) {
	for exitFlag := false; !exitFlag; {
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
			err = ScheduleMenu(config)
		case "3":
			err = ForecastMenu(config)
		case "4":
			//Exit()
			exitFlag = true
		default:
			fmt.Println("Invalid input")
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
