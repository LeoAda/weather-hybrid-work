package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Forecast struct {
	Time                string  `json:"time"`
	Temperature         float64 `json:"temperature_2m"`
	ApparentTemperature float64 `json:"apparent_temperature"`
	Precipitation       float64 `json:"precipitation"`
}

func (f Forecast) String() string {
	return fmt.Sprintf("Temperature: %f, Apparent Temperature: %f, Precipitation: %f, Time: %s", f.Temperature, f.ApparentTemperature, f.Precipitation, f.Time)
}

type ForecastList struct {
	Time                []string  `json:"time"`
	Temperature         []float64 `json:"temperature_2m"`
	ApparentTemperature []float64 `json:"apparent_temperature"`
	Precipitation       []float64 `json:"precipitation"`
}

type ForcastResponse struct {
	ForecastList ForecastList `json:"hourly"`
}

func GetForecast(Location Location) []Forecast {
	resp, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m,apparent_temperature,precipitation", Location.Latitude, Location.Longitude))
	if err != nil {
		return []Forecast{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Forecast{}
	}
	var forcastresponse ForcastResponse
	err = json.Unmarshal(body, &forcastresponse)
	if err != nil {
		return []Forecast{}
	}
	return forcastresponse.ForecastList.ConvertToListOfForecast()
}
func (f ForecastList) ConvertToListOfForecast() []Forecast {
	var forecast []Forecast
	for i := range f.Time {
		forecast = append(forecast, Forecast{
			Time:                f.Time[i],
			Temperature:         f.Temperature[i],
			ApparentTemperature: f.ApparentTemperature[i],
			Precipitation:       f.Precipitation[i],
		})
	}
	return forecast
}
