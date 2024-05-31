package jambonz

type SayVerb struct {
	Verb string `json:"verb"`
	Text string `json:"text"`
}

type GatherVerb struct {
	Verb       string   `json:"verb"`
	ActionHook string   `json:"actionHook"`
	Input      []string `json:"input,omitempty"`

	// TODO: can this be a rune?
	FinishOnKey string `json:"finishOnKey,omitempty"`
}

type CallStatus struct {
	CallSid    string `json:"call_sid"`
	CallStatus string `json:"call_status"`
}

type GatherResponse struct {
	Speech struct {
		Alternatives []struct {
			Transcript string
		}
	}
}

func Say(text string) *SayVerb {
	return &SayVerb{
		Verb: "say",
		Text: text,
	}
}

func Gather(actionHook string, input []string, finishOnKey rune) *GatherVerb {
	return &GatherVerb{
		Verb:        "gather",
		ActionHook:  actionHook,
		Input:       input,
		FinishOnKey: string(finishOnKey),
	}
}

func (gr *GatherResponse) GetTranscript() (string, bool) {
	if len(gr.Speech.Alternatives) == 0 {
		return "", false
	}

	transcript := gr.Speech.Alternatives[0].Transcript
	return transcript, true
}
