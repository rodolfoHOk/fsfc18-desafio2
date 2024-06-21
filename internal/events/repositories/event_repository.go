package repositories

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/rodolfoHOk/fsfc18-desafio2/internal/events/models"
)

const (
	EventNotFoundMessage    = "event not found"
	EventNotHasSpotsMessage = "event not has spots registered"
)

type EventRepository struct {
	events []models.Event
	spots  []models.Spot
}

func NewEventRepository() (*EventRepository, error) {
	data, err := loadData()
	if err != nil {
		return nil, err
	}

	return &EventRepository{
		events: data.Events,
		spots:  data.Spots,
	}, nil
}

func loadData() (*models.Data, error) {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var data models.Data
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *EventRepository) GetEvents() []models.Event {
	return r.events
}

func (r *EventRepository) GetEventByID(id int) (models.Event, error) {
	for _, event := range r.events {
		if event.ID == id {
			return event, nil
		}
	}
	return models.Event{}, errors.New(EventNotFoundMessage)
}

func (r *EventRepository) GetSpotsByEventID(eventId int) ([]models.Spot, error) {
	_, err := r.GetEventByID(eventId)
	if err != nil {
		return nil, err
	}

	var spots []models.Spot
	for _, spot := range r.spots {
		if spot.EventID == eventId {
			spots = append(spots, spot)
		}
	}

	if len(spots) == 0 {
		return nil, errors.New(EventNotHasSpotsMessage)
	}

	return spots, nil
}

func (r *EventRepository) ReserveSpots(eventId int, spotNames []string) ([]string, error) {
	_, err := r.GetEventByID(eventId)
	if err != nil {
		return nil, err
	}

	spots, err := r.GetSpotsByEventID(eventId)
	if err != nil {
		return nil, err
	}

	var notAvailable []string
	for _, spotName := range spotNames {
		idx := slices.IndexFunc(spots, func(s models.Spot) bool { return s.Name == spotName })
		if spots[idx].Status != "available" {
			notAvailable = append(notAvailable, spotName)
		}
	}
	if len(notAvailable) > 0 {
		notAvailableString := strings.Join(notAvailable, ", ")
		return nil, errors.New("Spots " + notAvailableString + " not available")
	}

	var reservedSpots []string
	for _, spotName := range spotNames {
		idx := slices.IndexFunc(spots, func(s models.Spot) bool { return s.Name == spotName })
		r.spots[idx].Status = "reserved"
		reservedSpots = append(reservedSpots, spotName)
	}

	return reservedSpots, nil
}
