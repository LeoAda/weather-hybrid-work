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
		startTime, endTime, err := askStartAndEndTimes()
		if err != nil {
			fmt.Println(err)
			continue
		}

		startHour, err := strconv.Atoi(startTime[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		startMin, err := strconv.Atoi(startTime[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		endHour, err := strconv.Atoi(endTime[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		endMin, err := strconv.Atoi(endTime[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		schedule = append(schedule, OutsideSchedule{
			StartHour: startHour,
			StartMin:  startMin,
			EndHour:   endHour,
			EndMin:    endMin,
		})

		if !askForMore() {
			break
		}
	}

	return schedule
}

func askStartAndEndTimes() ([]string, []string, error) {
	var startTime, endTime string
	fmt.Println("Enter start time (hh:mm): ")
	fmt.Scanln(&startTime)
	if !checkTimeFormat(startTime) {
		return nil, nil, fmt.Errorf("Invalid time format")
	}
	fmt.Println("Enter end time (hh:mm): ")
	fmt.Scanln(&endTime)
	if !checkTimeFormat(endTime) {
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

func checkTimeFormat(time string) bool {
	// check if time in format hh:mm with regex
	r, _ := regexp.Compile("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$")
	return r.MatchString(time)
}

func PrintSchedule(schedule []OutsideSchedule) {
	for _, s := range schedule {
		fmt.Println(s)
	}
}
