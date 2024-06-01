package calls

import (
	"strings"

	"github.com/hitecherik/emf-calls/internal/config"
	"github.com/hitecherik/emf-calls/pkg/fediverse"
)

type Handler interface {
	CanHandle(string, string) bool
	Handle(string, string) []interface{}
}

var handlers = []Handler{
	&GuestbookHandler{
		fedi: *fediverse.New(config.SnacApiKey, config.Url),
	},
	&BrightnessHandler{},
	LedHandler{},
	WeatherHandler{},
	EchoHandler{},
}

func Handle(text string, callSid string) []interface{} {
	lowercase := strings.ToLower(text)

	for _, handler := range handlers {
		if handler.CanHandle(lowercase, callSid) {
			return handler.Handle(text, callSid)
		}
	}

	return []interface{}{}
}
