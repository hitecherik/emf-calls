package calls

import "strings"

type Handler interface {
	CanHandle(string) bool
	Handle(string) []interface{}
}

var handlers = []Handler{WeatherHandler{}, LedHandler{}, EchoHandler{}}

func Handle(text string) []interface{} {
	for _, handler := range handlers {
		if handler.CanHandle(strings.ToLower(text)) {
			return handler.Handle(text)
		}
	}

	return []interface{}{}
}
