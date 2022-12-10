package main

import (
	"encoding/json"
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
	Locations      []Location `json:"results"`
	GenerationTime float64    `json:"generationtime_ms"`
}

func SearchLocationsList(query string) []Location {
	// Search for locations with Open Meteo Geocoding API
	resp, err := http.Get("https://geocoding-api.open-meteo.com/v1/search?name=" + query)
	if err != nil {
		return []Location{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Location{}
	}
	var locationResponse LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return []Location{}
	}
	return locationResponse.Locations
}
