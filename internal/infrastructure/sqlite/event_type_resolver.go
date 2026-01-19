package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
)

// EventTypeResolver is responsible for translating domain EventTypes
// into their corresponding database IDs (event_types.id).
//
// This is a pure infrastructure concern:
// - The domain does NOT know about IDs
// - The database does NOT know about domain enums
// - This adapter keeps both worlds decoupled
type EventTypeResolver struct {
	mapping map[domain.EventType]int64
}

// NewEventTypeResolver loads the event_types table into memory and
// builds a name -> id mapping.
//
// The application will fail fast if any required domain event type
// is missing from the database.
func NewEventTypeResolver(db *sql.DB) (*EventTypeResolver, error) {
	rows, err := db.Query(`SELECT id, name FROM event_types`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mapping := make(map[domain.EventType]int64)

	for rows.Next() {
		var id int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		mapping[domain.EventType(name)] = id
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Required domain event types that MUST exist in the database
	required := []domain.EventType{
		domain.SleepStart,
		domain.SleepEnd,
	}

	for _, t := range required {
		if _, ok := mapping[t]; !ok {
			return nil, fmt.Errorf("missing event type in database: %s", t)
		}
	}

	return &EventTypeResolver{mapping: mapping}, nil
}

// Resolve returns the database ID for a given domain EventType.
//
// It returns an error if the event type is unknown.
func (r *EventTypeResolver) Resolve(t domain.EventType) (int64, error) {
	id, ok := r.mapping[t]
	if !ok {
		return 0, fmt.Errorf("unknown event type: %s", t)
	}

	return id, nil
}
