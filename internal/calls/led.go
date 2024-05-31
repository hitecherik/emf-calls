package calls

import (
	"fmt"
	"strings"

	"github.com/hitecherik/emf-calls/internal/led"
	"github.com/hitecherik/emf-calls/pkg/jambonz"
)

type LedHandler struct{}

func (LedHandler) CanHandle(text string) bool {
	return strings.Contains(text, "tent")
}

func (LedHandler) Handle(text string) []interface{} {
	lowercase := strings.ToLower(text)

	for keyword, status := range led.StatusTranslation {
		if strings.Contains(lowercase, keyword) {
			led.SetStatus(status)
			return []interface{}{jambonz.Say(fmt.Sprintf("The tent has been turned %v!", keyword))}
		}
	}

	return []interface{}{jambonz.Say("Your desired tent state is not recognised.")}
}
