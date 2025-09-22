package httpsh

import (
	"encoding/json"
	"net/http"
	"time"

	"grep/2.18/domain"
	"grep/2.18/utils"
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

	if !utils.ValidStruct(resp) {
		h.logger.Error("empty event or empty field received")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"empty event or empty field received"})
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
	var resp domain.Event
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.logger.Error("decode error", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"invalid request body"})
		return
	}
	if !utils.ValidStruct(resp) {
		h.logger.Error("empty event or empty field received")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"empty event or empty field received"})
		return
	}

	if err := h.service.UpdateEvent(resp); err != nil {
		h.logger.Error("update event error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"update event error"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(domain.Response{Response: "event updated successfully"})
}

func (h *HandlerEvent) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var resp domain.Event
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		h.logger.Error("decode error", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"invalid request body"})
		return
	}
	if !utils.ValidStruct(resp) {
		h.logger.Error("empty event or empty field received")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"empty event or empty field received"})
		return
	}

	if err := h.service.DeleteEvent(resp); err != nil {
		h.logger.Error("delete event error")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"delete event error"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(domain.Response{Response: "event deleted successfully"})
}

func (h *HandlerEvent) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	var resp []domain.Event
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	userID := query.Get("user_id")
	userIdInt, err := utils.ParseUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"invalid user_id"})
		return
	}

	t, err := utils.ParseDate(query.Get("date"))
	if err != nil {
		if _, ok := err.(*time.ParseError); ok {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{"invalid date format, use YYYY-MM-DD"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"internal error"})
		return
	}

	resp, err = h.service.GetEventsForDay(userIdInt, t)
	if err != nil {
		h.logger.Error("get events error", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"user not found"})
		return
	}
	if len(resp) == 0 {
		h.logger.Error("dont have events for this date")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"dont have events for this date"})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *HandlerEvent) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	var resp []domain.Event
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	userID := query.Get("user_id")
	userIdInt, err := utils.ParseUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"invalid user_id"})
		return
	}

	t, err := utils.ParseDate(query.Get("date"))
	if err != nil {
		if _, ok := err.(*time.ParseError); ok {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{"invalid date format, use YYYY-MM-DD"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"internal error"})
		return
	}

	resp, err = h.service.GetEventsForWeek(userIdInt, t)
	if err != nil {
		h.logger.Error("get events error", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"user not found"})
		return
	}
	if len(resp) == 0 {
		h.logger.Error("dont have events for this date")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{"dont have events for this date"})
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *HandlerEvent) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
}
