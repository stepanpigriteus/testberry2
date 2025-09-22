package httpsh

import (
	"encoding/json"
	"net/http"

	"grep/2.18/domain"
)

type HandlerEvent struct {
	service domain.EventService
	logger  domain.Logger
}

func NewHandlerEvent(service domain.EventService, logger domain.Logger) *HandlerEvent {
	return &HandlerEvent{
		service: service,
		logger:  logger,
	}
}

func (h *HandlerEvent) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var resp domain.Event
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.logger.Error("decode error", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"invalid request body"})
		return
	}

	if resp == (domain.Event{}) {
		h.logger.Error("empty event received")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"empty event received"})
		return
	}

	if err := h.service.CreateEvent(resp); err != nil {
		h.logger.Error("create event error", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"create event error"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(domain.Response{Response: "event created successfully"})
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
