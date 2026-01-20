package contracts

import "github.com/dp-weasel/baby-sleep-tracker/internal/domain"

// EventReader defines read-only access to the event sequence.
type EventReader interface {
	// List returns events ordered chronologically.
	// Implementations must guarantee ordering.
	List(limit int) ([]domain.Event, error)

	// Last returns the most recent event in chronological order.
	// If there are no events, it returns (nil, nil).
	// Implementations must ensure ordering by event time, not insertion order.
	Last() (*domain.Event, error)
}
