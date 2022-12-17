package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Location        Location          `json:"location"`
	OutsideSchedule []OutsideSchedule `json:"outside_schedule"`
}

const configFile = "config.json"

func SaveConfig(config Config) error {
	file, err := os.Create(configFile) // creates if file doesn't exist
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(config, "", " ") // Parse object to JSON
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil { // Write JSON to file
		return err
	}
	return file.Sync()
}

func LoadConfig() (Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
