package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type OutsideSchedule struct {
	StartHour int `json:"start_hour"`
	StartMin  int `json:"start_minute"`
	EndHour   int `json:"end_hour"`
	EndMin    int `json:"end_minute"`
}

func (s OutsideSchedule) String() string {
	return fmt.Sprintf("%02d:%02d - %02d:%02d", s.StartHour, s.StartMin, s.EndHour, s.EndMin)
}

func AskDaySchedule() []OutsideSchedule {
	var schedule []OutsideSchedule

	for {
		startTime, endTime, err := askForStartAndEndTimes()
		if err != nil {
			fmt.Println(err)
			continue
		}
		StartHour, _ := strconv.Atoi(startTime[0])
		StartMin, _ := strconv.Atoi(startTime[1])
		EndHour, _ := strconv.Atoi(endTime[0])
		EndMin, _ := strconv.Atoi(endTime[1])
		schedule = append(schedule, OutsideSchedule{
			StartHour: StartHour,
			StartMin:  StartMin,
			EndHour:   EndHour,
			EndMin:    EndMin,
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
