package jambonz

type SayVerb struct {
	Verb string `json:"verb"`
	Text string `json:"text"`
}

type GatherVerb struct {
	Verb       string   `json:"verb"`
	ActionHook string   `json:"actionHook"`
	Input      []string `json:"input,omitempty"`
}

type PauseVerb struct {
	Verb   string `json:"verb"`
	Length int    `json:"length"`
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

func Gather(actionHook string, input []string) *GatherVerb {
	return &GatherVerb{
		Verb:       "gather",
		ActionHook: actionHook,
		Input:      input,
	}
}

func Pause(length int) *PauseVerb {
	return &PauseVerb{
		Verb:   "pause",
		Length: length,
	}
}

func (gr *GatherResponse) GetTranscript() (string, bool) {
	if len(gr.Speech.Alternatives) == 0 {
		return "", false
	}

	transcript := gr.Speech.Alternatives[0].Transcript
	return transcript, true
}
