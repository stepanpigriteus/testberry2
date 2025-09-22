package storage_test

import (
	"testing"
	"time"

	"grep/2.18/domain"
	"grep/2.18/storage"
)

func TestMemoryStorage(t *testing.T) {
	ms := storage.NewMemoryStorage()
	userID := 1

	date := time.Date(2025, 9, 22, 0, 0, 0, 0, time.UTC)

	event := domain.Event{
		UserID: userID,
		Date:   date,
		Event:  "Test Event",
	}

	if err := ms.CreateEvent(event); err != nil {
		t.Fatalf("CreateEvent failed: %v", err)
	}

	if err := ms.CreateEvent(event); err == nil {
		t.Fatalf("CreateEvent should fail on duplicate")
	}

	eventsDay, err := ms.GetEventsForDay(userID, date)
	if err != nil {
		t.Fatalf("GetEventsForDay failed: %v", err)
	}
	if len(eventsDay) != 1 {
		t.Fatalf("Expected 1 event, got %d", len(eventsDay))
	}

	updatedEvent := domain.Event{
		UserID: userID,
		Date:   date,
		Event:  "Updated Event",
	}
	if err := ms.UpdateEvent(updatedEvent); err != nil {
		t.Fatalf("UpdateEvent failed: %v", err)
	}

	eventsDay, _ = ms.GetEventsForDay(userID, date)
	if eventsDay[0].Event != "Updated Event" {
		t.Fatalf("UpdateEvent did not update event")
	}

	eventsWeek, err := ms.GetEventsForWeek(userID, date)
	if err != nil {
		t.Fatalf("GetEventsForWeek failed: %v", err)
	}
	if len(eventsWeek) != 1 {
		t.Fatalf("Expected 1 event for week, got %d", len(eventsWeek))
	}

	eventsMonth, err := ms.GetEventsForMonth(userID, date)
	if err != nil {
		t.Fatalf("GetEventsForMonth failed: %v", err)
	}
	if len(eventsMonth) != 1 {
		t.Fatalf("Expected 1 event for month, got %d", len(eventsMonth))
	}

	if err := ms.DeleteEvent(updatedEvent); err != nil {
		t.Fatalf("DeleteEvent failed: %v", err)
	}

	eventsDay, _ = ms.GetEventsForDay(userID, date)
	if len(eventsDay) != 0 {
		t.Fatalf("Expected 0 events after delete, got %d", len(eventsDay))
	}

	if err := ms.DeleteEvent(updatedEvent); err == nil {
		t.Fatalf("DeleteEvent should fail for non-existing event")
	}
}
