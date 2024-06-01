package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/hitecherik/emf-calls/internal/calls"
	"github.com/hitecherik/emf-calls/internal/config"
	"github.com/hitecherik/emf-calls/internal/led"
	"github.com/hitecherik/emf-calls/pkg/jambonz"
)

type contextCode uint8

const (
	callCompletedStatus  string = "completed"
	endOfCallPauseLength int    = 1

	requestBodyKey contextCode = iota
	callStatusKey
)

var activeSids sync.Map

func CallHandler(w http.ResponseWriter, r *http.Request) {
	activeSids.Store(r.Context().Value(callStatusKey).(*jambonz.CallStatus).CallSid, struct{}{})

	initialResponse := []interface{}{
		jambonz.Say("Welcome! Please say something and wait for a response."),
		jambonz.Gather(fmt.Sprintf("%v/talk", config.Url), []string{"speech"}, 0, ""),
	}

	rawResponse, err := json.Marshal(initialResponse)
	if err != nil {
		log.Printf("failed json marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rawResponse)
}

func TalkHandler(w http.ResponseWriter, r *http.Request) {
	rawRequest := r.Context().Value(requestBodyKey).([]byte)
	var request jambonz.GatherResponse

	if err := json.Unmarshal(rawRequest, &request); err != nil {
		log.Printf("failed JSON unmarshal: %v", err)
		return
	}

	sid := r.Context().Value(callStatusKey).(*jambonz.CallStatus).CallSid
	if _, ok := activeSids.Load(sid); !ok {
		log.Printf("access denied to sid %v", sid)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	transcript, _ := request.GetTranscript()
	response := append(
		calls.Handle(transcript, sid),
		jambonz.Pause(endOfCallPauseLength),
	)

	rawResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("failed json marshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(rawResponse)
}

func LogRequestBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Method == http.MethodPost {
			raw, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("reading request failed: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if len(raw) > 0 {
				log.Printf("received request: %v", string(raw))
				ctx = context.WithValue(ctx, requestBodyKey, raw)
			}
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func StoreCallStatusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Method == http.MethodPost {
			rawRequest := ctx.Value(requestBodyKey).([]byte)
			var status jambonz.CallStatus

			if err := json.Unmarshal(rawRequest, &status); err != nil {
				log.Printf("failed JSON unmarshal: %v", err)
				return
			}

			ctx = context.WithValue(ctx, callStatusKey, &status)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CallStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := r.Context().Value(callStatusKey).(*jambonz.CallStatus)
	if status.CallStatus == callCompletedStatus {
		activeSids.Delete(status.CallSid)
	}

	w.WriteHeader(http.StatusOK)
}

func LedColorHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(append([]byte{byte(led.GetMode())}, append([]byte(led.GetStatus()[:]), led.GetBrightness())...))
}

func LedColorConstantsHandler(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(led.StatusTranslation)
	if err != nil {
		log.Printf("could not marshal json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
