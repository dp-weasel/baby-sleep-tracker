package domain

// State represents the current state derived from sleep events.
type State string

const (
	StateEmpty  State = "EMPTY"  // No sufficient data to determine state
	StateAwake  State = "AWAKE"  // Baby is awake
	StateAsleep State = "ASLEEP" // Baby is sleeping
)

// StateFromEvents derives the state between two consecutive events.
// It assumes events are valid and ordered.
func StateFromEvents(prev EventType, next EventType) State {
	switch {
	case prev == SleepEnd && next == SleepStart:
		return StateAwake
	case prev == SleepStart && next == SleepEnd:
		return StateAsleep
	default:
		return StateEmpty
	}
}
