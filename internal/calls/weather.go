package calls

import (
	"log"
	"strings"

	"github.com/hitecherik/emf-calls/pkg/jambonz"
	"github.com/hitecherik/emf-calls/pkg/weather"
)

type WeatherHandler struct{}

func (WeatherHandler) CanHandle(text string, _ string) bool {
	return strings.Contains(text, "weather")
}

func (WeatherHandler) Handle(text string, _ string) []interface{} {
	currentWeather, err := weather.GetCurrentWeather()
	if err != nil {
		log.Printf("error getting weather: %v", err)
		return []interface{}{jambonz.Say("I'm sorry, I experienced an error getting the weather.")}
	}

	return []interface{}{jambonz.Say(currentWeather)}
}
