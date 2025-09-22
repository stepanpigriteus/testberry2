package serv

import (
	"log/slog"
	"time"

	"grep/2.18/domain"
)

type ImplEventServ struct {
	storage domain.Storage
	logger  domain.Logger
}

func NewServiceImpl(storage domain.Storage, logger domain.Logger) *ImplEventServ {
	return &ImplEventServ{
		storage: storage,
		logger:  logger,
	}
}

func (s *ImplEventServ) CreateEvent(event domain.Event) error {
	if err := s.storage.CreateEvent(event); err != nil {
		s.logger.Error("failed creating event (storage):", slog.Any("error", err))
		return err
	}
	return nil
}

func (s *ImplEventServ) UpdateEvent(event domain.Event) error {
	if err := s.storage.UpdateEvent(event); err != nil {
		s.logger.Error("failed updating event (storage):", slog.Any("error", err))
		return err
	}
	return nil
}

func (s *ImplEventServ) DeleteEvent(event domain.Event) error {
	if err := s.storage.DeleteEvent(event); err != nil {
		s.logger.Error("failed delete event (storage):", slog.Any("error", err))
		return err
	} else {
		return nil
	}
}

func (s *ImplEventServ) GetEventsForDay(userID int, date time.Time) ([]domain.Event, error) {
	if events, err := s.storage.GetEventsForDay(userID, date); err != nil {
		s.logger.Error("failed get event (storage):", slog.Any("error", err))
		return nil, err
	} else {
		return events, nil
	}
}

func (s *ImplEventServ) GetEventsForWeek(userID int, date time.Time) ([]domain.Event, error) {
	if events, err := s.storage.GetEventsForWeek(userID, date); err != nil {
		s.logger.Error("failed get event (storage):", slog.Any("error", err))
		return nil, err
	} else {
		return events, nil
	}
}

func (s *ImplEventServ) GetEventsForMonth(userID int, date time.Time) ([]domain.Event, error) {
	if events, err := s.storage.GetEventsForMonth(userID, date); err != nil {
		s.logger.Error("failed get event (storage):", slog.Any("error", err))
		return nil, err
	} else {
		return events, nil
	}
}
