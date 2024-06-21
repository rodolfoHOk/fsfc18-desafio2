package main

import (
	"net/http"

	"github.com/rodolfoHOk/fsfc18-desafio2/internal/events/api"
	"github.com/rodolfoHOk/fsfc18-desafio2/internal/events/repositories"
)

func main() {
	eventRepository, err := repositories.NewEventRepository()
	if err != nil {
		panic(err)
	}

	eventsHandler := api.NewEventsHandler(eventRepository)

	routesHandler := http.NewServeMux()
	routesHandler.HandleFunc("/events", eventsHandler.GetEvents)
	routesHandler.HandleFunc("/events/{eventID}", eventsHandler.GetEventByID)
	routesHandler.HandleFunc("/events/{eventID}/spots", eventsHandler.GetSpotsByEventID)
	routesHandler.HandleFunc("POST /events/{eventID}/reserve", eventsHandler.ReserveSpots)

	http.ListenAndServe(":8080", routesHandler)
}
