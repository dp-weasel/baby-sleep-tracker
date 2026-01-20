package query

import (
	"fmt"
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain/contracts"
)

// RecentCycle represents a summarized, UI-friendly view of a completed sleep/awake cycle.
type RecentCycle struct {
	From     time.Time
	To       time.Time
	Type     domain.State
	Duration string
	Note     string
}

// RecentCyclesQuery builds a list of recent completed cycles.
type RecentCyclesQuery struct {
	events contracts.EventReader
}

func NewRecentCyclesQuery(events contracts.EventReader) *RecentCyclesQuery {
	return &RecentCyclesQuery{events: events}
}

// Last returns the last N completed cycles, ordered from newest to oldest.
func (q *RecentCyclesQuery) Last(limit int) ([]RecentCycle, error) {
	// 1. Read all events (ordered ASC)
	events, err := q.events.List(0)
	if err != nil {
		return nil, err
	}

	cycles := []RecentCycle{}

	// 2. Build cycles from consecutive event pairs
	for i := 0; i+1 < len(events); i++ {
		prev := events[i]
		next := events[i+1]

		state := domain.StateFromEvents(prev.Type, next.Type)
		if state == domain.StateEmpty {
			continue
		}

		d := next.Timestamp.Sub(prev.Timestamp)
		cycles = append(cycles, RecentCycle{
			From:     prev.Timestamp,
			To:       next.Timestamp,
			Type:     state,
			Duration: formatDuration(d),
			Note:     next.Note,
		})
	}

	// 3. Keep only the last N cycles
	if limit > 0 && len(cycles) > limit {
		cycles = cycles[len(cycles)-limit:]
	}

	// 4. Reverse to newest-first
	reverseCycles(cycles)

	return cycles, nil
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60

	switch {
	case h > 0 && m > 0:
		return fmt.Sprintf("%dh %dm", h, m)
	case h > 0:
		return fmt.Sprintf("%dh", h)
	default:
		return fmt.Sprintf("%dm", m)
	}
}

func reverseCycles(c []RecentCycle) {
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
}
