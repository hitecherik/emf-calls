package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const endpoint = "https://api.open-meteo.com/v1/forecast?latitude=52.04&longitude=-2.38&current=temperature_2m,precipitation,wind_speed_10m,wind_gusts_10m"

type weatherData struct {
	Units struct {
		Temperature   string `json:"temperature_2m"`
		Precipitation string `json:"precipitation"`
		WindSpeed     string `json:"wind_speed_10m"`
		WindGusts     string `json:"wind_gusts_10m"`
	} `json:"current_units"`
	Data struct {
		Temperature   float64 `json:"temperature_2m"`
		Precipitation float64 `json:"precipitation"`
		WindSpeed     float64 `json:"wind_speed_10m"`
		WindGusts     float64 `json:"wind_gusts_10m"`
	} `json:"current"`
}

func GetCurrentWeather() (string, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var serialized weatherData
	if err := json.Unmarshal(respBody, &serialized); err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"The temperature is %v%v; there is %v%v of precipitation; the wind speed is %v%v; and there are wind gusts of %v%v.",
		serialized.Data.Temperature, serialized.Units.Temperature,
		serialized.Data.Precipitation, serialized.Units.Precipitation,
		serialized.Data.WindSpeed, serialized.Units.WindSpeed,
		serialized.Data.WindGusts, serialized.Units.WindGusts,
	), nil
}
