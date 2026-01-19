package application

import (
	"testing"
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
)

// inMemoryEventStore is a simple test double for EventStore
type inMemoryEventStore struct {
	events []domain.Event
}

func (s *inMemoryEventStore) Last() (*domain.Event, error) {
	if len(s.events) == 0 {
		return nil, nil
	}
	last := s.events[len(s.events)-1]
	return &last, nil
}

func (s *inMemoryEventStore) ExistsAt(ts time.Time) (bool, error) {
	for _, e := range s.events {
		if e.Timestamp.Equal(ts) {
			return true, nil
		}
	}
	return false, nil
}

func (s *inMemoryEventStore) Append(event domain.Event) error {
	s.events = append(s.events, event)
	return nil
}

func TestRegisterEvent_FirstEventMustBeWakeUp(t *testing.T) {
	store := &inMemoryEventStore{}
	service := &RegisterEventService{Store: store}

	err := service.Register(domain.SleepStart, time.Now(), "")
	if err != domain.ErrInvalidFirstEvent {
		t.Fatalf("expected ErrInvalidFirstEvent, got %v", err)
	}
}

func TestRegisterEvent_ValidSequence(t *testing.T) {
	store := &inMemoryEventStore{}
	service := &RegisterEventService{Store: store}

	t1 := time.Date(2026, 1, 10, 7, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, 1, 10, 8, 0, 0, 0, time.UTC)

	if err := service.Register(domain.SleepEnd, t1, ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := service.Register(domain.SleepStart, t2, ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
