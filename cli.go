package main

import "fmt"

func LocationMenu(config Config) error {
	location, err := ChooseLocation()
	if err != nil {
		return fmt.Errorf("Error choosing location: %v", err)
	} else {
		config.Location = location
		return config.Save()
	}
}
