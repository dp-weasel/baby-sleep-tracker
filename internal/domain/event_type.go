package domain

// EventType represents the semantic type of an event in the domain.
type EventType string

const (
	SleepStart EventType = "SUEÑO_INICIO"
	SleepEnd   EventType = "SUEÑO_FIN"
)
