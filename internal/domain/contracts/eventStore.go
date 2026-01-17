package contracts

import (
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
)

// EventStore defines write operations over the event sequence.
// Implementations must guarantee consistency and ordering.
type EventStore interface {
	// Last returns the most recently registered event.
	// If no events exist, it returns (nil, nil).
	Last() (*domain.Event, error)

	// ExistsAt checks whether an event already exists at the given timestamp.
	ExistsAt(ts time.Time) (bool, error)

	// Append persists a new event to the sequence.
	Append(event domain.Event) error
}
