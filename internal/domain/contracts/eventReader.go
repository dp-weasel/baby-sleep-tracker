package contracts

import "github.com/dp-weasel/baby-sleep-tracker/internal/domain"

// EventReader defines read-only access to the event sequence.
type EventReader interface {
	// List returns events ordered chronologically.
	// Implementations must guarantee ordering.
	List(limit int) ([]domain.Event, error)
}
