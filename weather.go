package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ForecastList struct {
	Time                []string  `json:"time"`
	Temperature         []float64 `json:"temperature_2m"`
	ApparentTemperature []float64 `json:"apparent_temperature"`
	Precipitation       []float64 `json:"precipitation"`
}

/*
	type Forecast struct {
		Temperature         float64 `json:"temperature_2m"`
		ApparentTemperature float64 `json:"apparent_temperature"`
		Precipitation       float64 `json:"precipitation"`
	}
*/
func (f ForecastList) String() string {
	return fmt.Sprintf("Temperature: %f, Apparent Temperature: %f, Precipitation: %f, Time: %s", f.Temperature, f.ApparentTemperature, f.Precipitation, f.Time)
}

type ForcastResponse struct {
	ForecastList ForecastList `json:"hourly"`
}

func GetForecast(Location Location) ForecastList {
	resp, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m,apparent_temperature,precipitation", Location.Latitude, Location.Longitude))
	if err != nil {
		return ForecastList{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ForecastList{}
	}
	var forecastResponse ForcastResponse
	err = json.Unmarshal(body, &forecastResponse)
	if err != nil {
		return ForecastList{}
	}
	return forecastResponse.ForecastList
}
