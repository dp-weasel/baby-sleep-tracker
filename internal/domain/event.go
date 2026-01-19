package domain

import (
	"time"
)

// Event represents an immutable fact that occurred at a specific point in time.
type Event struct {
	Type      EventType
	Timestamp time.Time
	Note      string
}

// NewEvent creates a new domain event.
// Ordering and sequence validation are handled by the application layer.
func NewEvent(eventType EventType, ts time.Time, note string) Event {
	return Event{
		Type:      eventType,
		Timestamp: ts,
		Note:      note,
	}
}
