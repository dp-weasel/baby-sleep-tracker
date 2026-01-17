package application

import (
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain/contracts"
)

// QueryPeriodsService orchestrates the use case of deriving periods from events.
type QueryPeriodsService struct {
	Reader contracts.EventReader
}

// Query derives periods from the ordered event sequence.
func (s *QueryPeriodsService) Query(limit int) ([]domain.Period, error) {
	events, err := s.Reader.List(limit)
	if err != nil {
		return nil, err
	}

	if len(events) < 2 {
		return nil, domain.ErrInsufficientEvents
	}

	periods := make([]domain.Period, 0, len(events)-1)

	for i := 1; i < len(events); i++ {
		p := domain.DerivePeriod(events[i-1], events[i])
		periods = append(periods, p)
	}

	return periods, nil
}
