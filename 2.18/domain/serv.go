package domain

import "time"

type EventService interface {
	CreateEvent(event Event) error
	UpdateEvent(event Event) error
	DeleteEvent(event Event) error
	GetEventsForDay(userID int, date time.Time) ([]Event, error)
	GetEventsForWeek(userID int, weekStart time.Time) ([]Event, error)
	GetEventsForMonth(userID int, month time.Time) ([]Event, error)
}
