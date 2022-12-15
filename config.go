package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Location        Location          `json:"location"`
	OutsideSchedule []OutsideSchedule `json:"outside_schedule"`
}

func SaveConfig(config Config) error {
	//Save config to config.json
	file, err := os.Create("config.json") // creates if file doesn't exist
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(config, "", " ") // Parse object to JSON
	if err != nil {
		return err
	}
	_, err = file.Write(data) // Write JSON to file
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() (Config, error) {
	// Load config from config.json
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
