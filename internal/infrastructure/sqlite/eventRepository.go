package sqlite

import (
	"database/sql"
	"time"

	"github.com/dp-weasel/baby-sleep-tracker/internal/domain"
	"github.com/dp-weasel/baby-sleep-tracker/internal/domain/contracts"
)

// EventRepository implements EventStore and EventReader using SQLite.
type EventRepository struct {
	db       *sql.DB
	resolver *EventTypeResolver
}

func NewEventRepository(db *sql.DB, resolver *EventTypeResolver) *EventRepository {
	return &EventRepository{
		db:       db,
		resolver: resolver,
	}
}

// Ensure EventRepository implements required interfaces
var _ contracts.EventStore = (*EventRepository)(nil)
var _ contracts.EventReader = (*EventRepository)(nil)

func (r *EventRepository) Last() (*domain.Event, error) {
	row := r.db.QueryRow(`
		SELECT et.name, al.event_time, al.note
		FROM activity_logs al
		JOIN event_types et ON et.id = al.event_type_id
		ORDER BY al.event_time DESC
		LIMIT 1
	`)

	var typeName string
	var eventTime string
	var note sql.NullString

	err := row.Scan(&typeName, &eventTime, &note)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	ts, err := time.Parse(time.RFC3339, eventTime)
	if err != nil {
		return nil, err
	}

	return &domain.Event{
		Type:      domain.EventType(typeName),
		Timestamp: ts,
		Note:      note.String,
	}, nil
}

func (r *EventRepository) ExistsAt(ts time.Time) (bool, error) {
	row := r.db.QueryRow(`
		SELECT COUNT(1)
		FROM activity_logs
		WHERE event_time = ?
		`, ts.Format(time.RFC3339))

	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *EventRepository) Append(event domain.Event) error {
	eventTypeID, err := r.resolver.Resolve(event.Type)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		INSERT INTO activity_logs (event_type_id, event_time, note)
		VALUES (?, ?, ?)
		`, eventTypeID, event.Timestamp.Format(time.RFC3339), event.Note)

	return err
}

func (r *EventRepository) List(limit int) ([]domain.Event, error) {
	query := `
		SELECT et.name, al.event_time, al.note
		FROM activity_logs al
		JOIN event_types et ON et.id = al.event_type_id
		ORDER BY al.event_time ASC
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
		var typeName string
		var eventTime string
		var note sql.NullString

		if err := rows.Scan(&typeName, &eventTime, &note); err != nil {
			return nil, err
		}

		ts, err := time.Parse(time.RFC3339, eventTime)
		if err != nil {
			return nil, err
		}

		events = append(events, domain.Event{
			Type:      domain.EventType(typeName),
			Timestamp: ts,
			Note:      note.String,
		})
	}

	return events, rows.Err()
}
