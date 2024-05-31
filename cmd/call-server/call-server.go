package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hitecherik/emf-calls/internal/config"
	"github.com/hitecherik/emf-calls/internal/server"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(server.LogRequestBodyMiddleware)
	r.Use(server.StoreCallStatusMiddleware)

	authed := chi.NewRouter()
	authed.Use(middleware.BasicAuth(config.Hostname, config.BasicAuthCredentials))
	authed.Post("/call", server.CallHandler)
	authed.Post("/status", server.CallStatusHandler)

	r.Mount("/authed", authed)
	r.Post("/talk", server.TalkHandler)
	r.Get("/led", server.LedColorHandler)
	r.Get("/led/constants", server.LedColorConstantsHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), r))
}
