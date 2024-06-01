package calls

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/hitecherik/emf-calls/internal/config"
	"github.com/hitecherik/emf-calls/internal/led"
	"github.com/hitecherik/emf-calls/pkg/jambonz"
)

var (
	numberRe      = regexp.MustCompile(`\b(1?\d\d?|2[0-4]\d|25[0-5])\b`)
	colorChangeRe = regexp.MustCompile(`(tent|colou?r)`)
)

type LedHandler struct{}

func (LedHandler) CanHandle(text string, _ string) bool {
	return colorChangeRe.MatchString(text)
}

func (LedHandler) Handle(text string, _ string) []interface{} {
	lowercase := strings.ToLower(text)

	for keyword, status := range led.StatusTranslation {
		if strings.Contains(lowercase, keyword) {
			led.SetStatus(status)
			return []interface{}{jambonz.Say(fmt.Sprintf("The LEDs have been changed to %v!", keyword))}
		}
	}

	return []interface{}{jambonz.Say("Your desired LED state is not recognised.")}
}

type BrightnessHandler struct {
	states sync.Map
}

func (bh *BrightnessHandler) CanHandle(text string, callSid string) bool {
	_, ok := bh.states.Load(callSid)
	return ok || strings.Contains(text, "brightness")
}

func (bh *BrightnessHandler) Handle(text string, callSid string) []interface{} {
	if _, loaded := bh.states.LoadOrStore(callSid, struct{}{}); !loaded {
		return []interface{}{
			jambonz.Say("Please select a brightness between 0 and 255."),
			jambonz.Gather(fmt.Sprintf("%v/talk", config.Url), []string{"speech"}),
		}
	}

	bh.states.Delete(callSid)

	rawBrightness := numberRe.FindString(text)
	if rawBrightness == "" {
		return []interface{}{
			jambonz.Say("Your brightness has not been recognised."),
		}
	}

	// Given the regex above has matched, this won't fail
	brightness, _ := strconv.ParseUint(rawBrightness, 10, 8)
	led.SetBrightness(byte(brightness))

	return []interface{}{
		jambonz.Say(fmt.Sprintf("The LED brightness has been set to %v.", rawBrightness)),
	}
}
