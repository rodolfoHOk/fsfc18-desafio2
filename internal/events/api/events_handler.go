package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rodolfoHOk/fsfc18-desafio2/internal/events/repositories"
)

type EventsHandler struct {
	repository *repositories.EventRepository
}

func NewEventsHandler(repository *repositories.EventRepository) *EventsHandler {
	return &EventsHandler{
		repository: repository,
	}
}

func (h *EventsHandler) GetEvents(writer http.ResponseWriter, request *http.Request) {
	responseBody := h.repository.GetEvents()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(responseBody)
}

func (h *EventsHandler) GetEventByID(writer http.ResponseWriter, request *http.Request) {
	eventID, err := strconv.Atoi(request.PathValue("eventID"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, err := h.repository.GetEventByID(int(eventID))
	if err != nil {
		if err.Error() == repositories.EventNotFoundMessage {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(responseBody)
}

func (h *EventsHandler) GetSpotsByEventID(writer http.ResponseWriter, request *http.Request) {
	eventID, err := strconv.Atoi(request.PathValue("eventID"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, err := h.repository.GetSpotsByEventID(eventID)
	if err != nil {
		if err.Error() == repositories.EventNotFoundMessage || err.Error() == repositories.EventNotHasSpotsMessage {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(responseBody)
}

type ReserveSpotsDTO struct {
	Spots []string `json:"spots"`
}

func (h *EventsHandler) ReserveSpots(writer http.ResponseWriter, request *http.Request) {
	eventID, err := strconv.Atoi(request.PathValue("eventID"))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var requestBody ReserveSpotsDTO
	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	responseBody, err := h.repository.ReserveSpots(eventID, requestBody.Spots)
	if err != nil {
		if err.Error() == repositories.EventNotFoundMessage || err.Error() == repositories.EventNotHasSpotsMessage {
			http.Error(writer, err.Error(), http.StatusNotFound)
			return
		} else if strings.HasSuffix(err.Error(), "not available") {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(responseBody)
}
