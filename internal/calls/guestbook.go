package calls

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hitecherik/emf-calls/internal/config"
	"github.com/hitecherik/emf-calls/pkg/fediverse"
	"github.com/hitecherik/emf-calls/pkg/jambonz"
)

type GuestbookHandler struct {
	states sync.Map
	fedi   fediverse.Fediverse
}

func (gb *GuestbookHandler) CanHandle(text string, callSid string) bool {
	_, ok := gb.states.Load(callSid)
	return ok || strings.Contains(text, "post")
}

func (gb *GuestbookHandler) Handle(text string, callSid string) []interface{} {
	if _, loaded := gb.states.LoadOrStore(callSid, struct{}{}); !loaded {
		return []interface{}{
			jambonz.Say("Please dictate your post."),
			jambonz.Gather(fmt.Sprintf("%v/talk", config.Url), []string{"speech"}, 0, ""),
		}
	}

	gb.states.Delete(callSid)

	if err := gb.fedi.Post(text); err != nil {
		log.Printf(`posting message "%v" failed: %v`, text, err)
		return []interface{}{
			jambonz.Say("There was an error posting your message."),
		}
	}

	return []interface{}{
		jambonz.Say("Your post has been successfully created."),
	}
}
