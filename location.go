package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	City       string  `json:"name"`
	Country    string  `json:"country"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Complement string  `json:"admin1"`
}

func (l Location) String() string {
	return l.City + ", " + l.Complement + ", " + l.Country
}

type LocationResponse struct {
	Locations []Location `json:"results"`
}

func GetLocationsList(query string) ([]Location, error) {
	// Search for locations with Open Meteo Geocoding API
	resp, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?name=" + query)
	if err != nil {
		return []Location{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Location{}, err
	}
	var locationResponse LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return []Location{}, err
	}
	return locationResponse.Locations, nil
}

func ChooseLocation() (Location, error) {
	// Display locations and ask user to choose one, return the chosen location
	fmt.Print("Type city name: ")
	var cityName string
	if _, err := fmt.Scan(&cityName); err != nil {
		return Location{}, fmt.Errorf("Error reading city name: %v", err)
	}

	locations, err := GetLocationsList(cityName)
	if err != nil {
		return Location{}, fmt.Errorf("Error getting locations list: %v", err)
	}

	for index, element := range locations {
		fmt.Printf("%d. %s\n", index, element)
	}

	var choice int
	fmt.Print("Choose a location: ")
	if _, err := fmt.Scan(&choice); err != nil {
		return Location{}, fmt.Errorf("Error reading choice: %v", err)
	}

	if choice < 0 || choice >= len(locations) {
		return Location{}, fmt.Errorf("Invalid choice")
	}

	return locations[choice], nil
}
