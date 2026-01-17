package sqlite

import (
	"database/sql"
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain/contracts"
)

// EventRepository implements EventStore and EventReader using SQLite.
type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

// Ensure EventRepository implements required interfaces
var _ contracts.EventStore = (*EventRepository)(nil)
var _ contracts.EventReader = (*EventRepository)(nil)

func (r *EventRepository) Last() (*domain.Event, error) {
	row := r.db.QueryRow(`
		SELECT type, timestamp
		FROM events
		ORDER BY timestamp DESC
		LIMIT 1
	`)

	var e domain.Event
	var ts int64

	err := row.Scan(&e.Type, &ts)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	e.Timestamp = time.Unix(ts, 0)
	return &e, nil
}

func (r *EventRepository) ExistsAt(ts time.Time) (bool, error) {
	row := r.db.QueryRow(`
		SELECT COUNT(1)
		FROM events
		WHERE timestamp = ?
	`, ts.Unix())

	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *EventRepository) Append(event domain.Event) error {
	_, err := r.db.Exec(`
		INSERT INTO events (type, timestamp)
		VALUES (?, ?)
	`, event.Type, event.Timestamp.Unix())
	return err
}

func (r *EventRepository) List(limit int) ([]domain.Event, error) {
	query := `
		SELECT type, timestamp
		FROM events
		ORDER BY timestamp ASC
	`

	args := []any{}
	if limit > 0 {
		query += " LIMIT ?"
		args = append(args, limit)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := []domain.Event{}

	for rows.Next() {
		var e domain.Event
		var ts int64

		if err := rows.Scan(&e.Type, &ts); err != nil {
			return nil, err
		}

		e.Timestamp = time.Unix(ts, 0)
		events = append(events, e)
	}

	return events, nil
}
