package application

import (
	"testing"
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
)

// inMemoryEventReader is a simple test double for EventReader
type inMemoryEventReader struct {
	events []domain.Event
}

func (r *inMemoryEventReader) List(limit int) ([]domain.Event, error) {
	if limit > 0 && limit < len(r.events) {
		return r.events[:limit], nil
	}
	return r.events, nil
}

func TestQueryPeriods_DerivesCorrectPeriods(t *testing.T) {
	t1 := time.Date(2026, 1, 10, 7, 0, 0, 0, time.UTC)
	t2 := time.Date(2026, 1, 10, 8, 0, 0, 0, time.UTC)
	t3 := time.Date(2026, 1, 10, 10, 0, 0, 0, time.UTC)

	reader := &inMemoryEventReader{events: []domain.Event{
		{Type: domain.SleepEnd, Timestamp: t1},
		{Type: domain.SleepStart, Timestamp: t2},
		{Type: domain.SleepEnd, Timestamp: t3},
	}}

	service := &QueryPeriodsService{Reader: reader}

	periods, err := service.Query(0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(periods) != 2 {
		t.Fatalf("expected 2 periods, got %d", len(periods))
	}

	if periods[0].Type != domain.Awake {
		t.Fatalf("expected first period to be Awake")
	}

	if periods[1].Type != domain.Sleeping {
		t.Fatalf("expected second period to be Sleeping")
	}
}
