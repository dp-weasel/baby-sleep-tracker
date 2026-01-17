package application

import (
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain/contracts"
)

// RegisterEventService orchestrates the use case of registering a new event.
type RegisterEventService struct {
	Store contracts.EventStore
}

// Register registers a new domain event enforcing domain rules.
func (s *RegisterEventService) Register(eventType domain.EventType, ts time.Time) error {
	event := domain.NewEvent(eventType, ts)

	last, err := s.Store.Last()
	if err != nil {
		return err
	}

	// First event rule
	if last == nil {
		if event.Type != domain.SleepEnd {
			return domain.ErrInvalidFirstEvent
		}
		return s.Store.Append(event)
	}

	// Timestamp ordering
	if !event.Timestamp.After(last.Timestamp) {
		return domain.ErrOutOfOrder
	}

	// Consecutive type validation
	if event.Type == last.Type {
		return domain.ErrConsecutiveSameType
	}

	// Duplicate timestamp validation
	exists, err := s.Store.ExistsAt(event.Timestamp)
	if err != nil {
		return err
	}
	if exists {
		return domain.ErrSameTimestamp
	}

	return s.Store.Append(event)
}
