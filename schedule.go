package main

import (
	"fmt"
	"regexp"
	"strings"
)

type OutsideSchedule struct {
	StartHour string `json:"start_hour"`
	StartMin  string `json:"start_minute"`
	EndHour   string `json:"end_hour"`
	EndMin    string `json:"end_minute"`
}

func (s OutsideSchedule) String() string {
	return fmt.Sprintf("Start: %s:%s, End: %s:%s", s.StartHour, s.StartMin, s.EndHour, s.EndMin)
}

func AskDaySchedule() []OutsideSchedule {
	var schedule []OutsideSchedule

	for {
		startTime, endTime, err := askForStartAndEndTimes()
		if err != nil {
			fmt.Println(err)
			continue
		}

		schedule = append(schedule, OutsideSchedule{
			StartHour: startTime[0],
			StartMin:  startTime[1],
			EndHour:   endTime[0],
			EndMin:    endTime[1],
		})

		if !askForMore() {
			break
		}
	}

	return schedule
}

func askForStartAndEndTimes() ([]string, []string, error) {
	var startTime, endTime string
	fmt.Println("Enter start time (hh:mm): ")
	fmt.Scanln(&startTime)
	if !CheckTimeFormat(startTime) {
		return nil, nil, fmt.Errorf("Invalid time format")
	}
	fmt.Println("Enter end time (hh:mm): ")
	fmt.Scanln(&endTime)
	if !CheckTimeFormat(endTime) {
		return nil, nil, fmt.Errorf("Invalid time format")
	}

	startTimeList := strings.Split(startTime, ":")
	endTimeList := strings.Split(endTime, ":")
	return startTimeList, endTimeList, nil
}

func askForMore() bool {
	var more string
	fmt.Println("More? (y/n): ")
	fmt.Scanln(&more)
	return more == "y"
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
