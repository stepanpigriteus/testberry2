package httpsh

import (
	"net/http"

	"grep/2.18/domain"
)

type HandlerEvent struct {
	service domain.EventService
}

func NewHandlerEvent(service domain.EventService) *HandlerEvent {
	return &HandlerEvent{
		service: service,
	}
}

func (h *HandlerEvent) CreateEvent(w http.ResponseWriter, r *http.Request) {
}

func (h *HandlerEvent) UpdateEvent(w http.ResponseWriter, r *http.Request) {
}

func (h *HandlerEvent) DeleteEvent(w http.ResponseWriter, r *http.Request) {
}

func (h *HandlerEvent) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
}

func (h *HandlerEvent) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
}

func (h *HandlerEvent) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
}
