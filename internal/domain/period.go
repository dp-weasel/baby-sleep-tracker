package domain

import "time"

// PeriodType represents the semantic meaning of a derived period.
type PeriodType string

const (
	Awake    PeriodType = "DESPIERTO"
	Sleeping PeriodType = "DURMIENDO"
)

// Period represents a derived time span between two consecutive events.
// Periods are never persisted.
type Period struct {
	From     time.Time
	To       time.Time
	Type     PeriodType
	Duration time.Duration
}

// DerivePeriod derives a Period from two consecutive events.
// It assumes events are already validated and ordered.
func DerivePeriod(prev, next Event) Period {
	var pType PeriodType

	switch prev.Type {
	case SleepEnd:
		pType = Awake
	case SleepStart:
		pType = Sleeping
	}

	duration := next.Timestamp.Sub(prev.Timestamp)

	return Period{
		From:     prev.Timestamp,
		To:       next.Timestamp,
		Type:     pType,
		Duration: duration,
	}
}
