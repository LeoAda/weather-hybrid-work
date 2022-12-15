package main

import (
	"fmt"
	"regexp"
)

type OutsideSchedule struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (s OutsideSchedule) String() string {
	return fmt.Sprintf("Start time: %s, End time: %s", s.StartTime, s.EndTime)
}

func AskDaySchedule() []OutsideSchedule {
	var schedule []OutsideSchedule
	for {
		var start_time, end_time string
		fmt.Println("Enter start time (hh:mm): ")
		fmt.Scanln(&start_time)
		if !CheckTimeFormat(start_time) {
			fmt.Println("Invalid time format")
			continue
		}
		fmt.Println("Enter end time (hh:mm): ")
		fmt.Scanln(&end_time)
		if !CheckTimeFormat(end_time) {
			fmt.Println("Invalid time format")
			continue
		}
		schedule = append(schedule, OutsideSchedule{
			StartTime: start_time,
			EndTime:   end_time,
		})
		var more string
		fmt.Println("More? (y/n): ")
		fmt.Scanln(&more)
		if more == "n" {
			break
		}
	}

	return schedule
}

func CheckTimeFormat(time string) bool {
	// check if time in format hh:mm with regex
	r, _ := regexp.Compile("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$")
	return r.MatchString(time)
}

func PrintSchedule(schedule []OutsideSchedule) {
	for _, s := range schedule {
		fmt.Println(s)
	}
}
