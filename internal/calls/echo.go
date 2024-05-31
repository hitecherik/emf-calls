package calls

import (
	"fmt"

	"github.com/hitecherik/emf-calls/pkg/jambonz"
)

type EchoHandler struct{}

func (EchoHandler) CanHandle(_ string) bool {
	return true
}

func (EchoHandler) Handle(text string) []interface{} {
	if len(text) > 0 {
		return []interface{}{jambonz.Say(fmt.Sprintf("This is what you said: %v", text))}
	}

	return []interface{}{jambonz.Say("It doesn't sound like you said anything.")}
}
