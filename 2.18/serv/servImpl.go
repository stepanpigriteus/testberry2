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
	return nil
}

func (s *ImplEventServ) DeleteEvent(userID int, eventDate time.Time) error {
	return nil
}

func (s *ImplEventServ) GetEventsForDay(userID int, date time.Time) ([]domain.Event, error) {

	var events []domain.Event
	return events, nil
}

func (s *ImplEventServ) GetEventsForWeek(userID int, date time.Time) ([]domain.Event, error) {
	var events []domain.Event
	return events, nil
}

func (s *ImplEventServ) GetEventsForMonth(userID int, date time.Time) ([]domain.Event, error) {
	var events []domain.Event
	return events, nil
}
