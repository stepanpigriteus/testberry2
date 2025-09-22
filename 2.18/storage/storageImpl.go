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
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.events[event.UserID]
	if !ok {
		return fmt.Errorf("no events found for user %d", event.UserID)
	}
	updated := false
	for i, ev := range val {
		if ev.Date.Equal(event.Date) || ev.Event == event.Event {
			val[i] = event
			updated = true
			break
		}
	}
	if !updated {
		return fmt.Errorf("event not found for user %d on date %s", event.UserID, event.Date.Format("2006-01-02"))
	}
	m.events[event.UserID] = val
	return nil
}

func (m *MemoryStorage) DeleteEvent(event domain.Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	val, ok := m.events[event.UserID]
	if !ok {
		return fmt.Errorf("user not found, id %d", event.UserID)
	}
	found := false
	for i, r := range val {
		if r == event {
			val = append(val[:i], val[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("event not found for user %d", event.UserID)
	}
	if len(val) == 0 {
		delete(m.events, event.UserID)
	} else {
		m.events[event.UserID] = val
	}
	return nil
}

func (m *MemoryStorage) GetEventsForDay(userID int, date time.Time) ([]domain.Event, error) {
	 m.mu.RLock()
    defer m.mu.RUnlock()
	var events []domain.Event
	val, ok := m.events[userID]
	if !ok {
		return nil, fmt.Errorf("user not found, id %d", userID)
	}
	for _, r := range val {
		if r.Date.Year() == date.Year() &&
			r.Date.Month() == date.Month() &&
			r.Date.Day() == date.Day() {
			events = append(events, r)
		}
	}

	return events, nil
}

func (m *MemoryStorage) GetEventsForWeek(userID int, date time.Time) ([]domain.Event, error) {
	 m.mu.RLock()
    defer m.mu.RUnlock()
	var events []domain.Event
	val, ok := m.events[userID]
	if !ok {
		return nil, fmt.Errorf("no events found for user %d", userID)
	}
	startOfWeek := date.Truncate(24 * time.Hour)
	for startOfWeek.Weekday() != time.Monday {
		startOfWeek = startOfWeek.AddDate(0, 0, -1)
	}
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	for _, r := range val {
		if !r.Date.Before(startOfWeek) && r.Date.Before(endOfWeek) {
			events = append(events, r)
		}
	}
	return events, nil
}

func (m *MemoryStorage) GetEventsForMonth(userID int, date time.Time) ([]domain.Event, error) {
	 m.mu.RLock()
    defer m.mu.RUnlock()
	var events []domain.Event
	val, ok := m.events[userID]
	if !ok {
		return nil, fmt.Errorf("no events found for user %d", userID)
	}
	startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	for _, r := range val {
		if !r.Date.Before(startOfMonth) && r.Date.Before(endOfMonth) {
			events = append(events, r)
		}
	}
	return events, nil
}
