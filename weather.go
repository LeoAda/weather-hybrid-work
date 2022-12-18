package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Forecast struct {
	Date                string  `json:"time"`
	Hour                int     `json:"hour"`
	Temperature         float64 `json:"temperature_2m"`
	ApparentTemperature float64 `json:"apparent_temperature"`
	Precipitation       float64 `json:"precipitation"`
}

func (f Forecast) String() string {
	return fmt.Sprintf("%s %d:00: %.1f°C, %.1f°C, %.1fmm", f.Date, f.Hour, f.Temperature, f.ApparentTemperature, f.Precipitation)
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

func GetWeekForecast(Location Location, StartDay time.Time) []Forecast {
	EndTime := StartDay.AddDate(0, 0, 4)
	resp, err := http.Get(fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m,apparent_temperature,precipitation&start_date=%s&end_date=%s", Location.Latitude, Location.Longitude, StartDay.Format("2006-01-02"), EndTime.Format("2006-01-02")))
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
		hour, err := time.Parse("2006-02-15T15:04", f.Time[i])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(hour)
		forecast = append(forecast, Forecast{
			Date:                f.Time[i],
			Hour:                hour.Hour(),
			Temperature:         f.Temperature[i],
			ApparentTemperature: f.ApparentTemperature[i],
			Precipitation:       f.Precipitation[i],
		})
	}
	return forecast
}
