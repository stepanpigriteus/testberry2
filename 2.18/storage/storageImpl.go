package storage

import (
	"fmt"
	"sync"
	"time"

	"grep/2.18/domain"
)

type MemoryStorage struct {
	events map[int][]domain.Event
	mu     sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		events: make(map[int][]domain.Event),
	}
}

func (m *MemoryStorage) CreateEvent(event domain.Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	fmt.Println(event.Event)
		for _, r := range m.events[event.UserID] {
		if r.Date == event.Date && r.Event == event.Event {
			return fmt.Errorf("the event exists: %s", event.Event)
		}
	}
	m.events[event.UserID] = append(m.events[event.UserID], event)
	fmt.Println(m.events)
	return nil
}

func (m *MemoryStorage) UpdateEvent(event domain.Event) error {
	return nil
}

func (m *MemoryStorage) DeleteEvent(userID int, eventDate time.Time) error {
	return nil
}

func (m *MemoryStorage) GetEventsForDay(userID int, date time.Time) ([]domain.Event, error) {
	var events []domain.Event
	return events, nil
}

func (m *MemoryStorage) GetEventsForWeek(userID int, weekStart time.Time) ([]domain.Event, error) {
	var events []domain.Event
	return events, nil
}

func (m *MemoryStorage) GetEventsForMonth(userID int, month time.Time) ([]domain.Event, error) {
	var events []domain.Event
	return events, nil
}
