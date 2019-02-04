package event

import (
	"github.com/lib/pq"
	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/database"
)

type Store struct {
	db *database.Database
}

func NewStore(db *database.Database) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Save(events []Event) (err error) {
	txn, err := s.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = txn.Rollback()
			return
		}
		err = txn.Commit()
	}()

	stmt, err := txn.Prepare(pq.CopyIn("events", "id", "aggregate_id", "kind", "version", "fired_at", "data"))
	if err != nil {
		return
	}

	for _, e := range events {
		_, err = stmt.Exec(e.ID(), e.AggregateID(), e.Kind(), e.Version(), e.FiredAt(), string(e.Data()))
		if err != nil {
			return err
		}
	}

	if _, err = stmt.Exec(); err != nil {
		return
	}

	if err = stmt.Close(); err != nil {
		return
	}

	return
}

func (s *Store) Load(aggregateID id.ID, fromVersion uint64) ([]Event, error) {
	rows, err := s.db.Query(
		"SELECT id, kind, version, fired_at, data FROM events WHERE aggregate_id = $1 AND version > $2 ORDER BY fired_at DESC",
		aggregateID,
		fromVersion,
	)
	if err != nil {
		return nil, err
	}

	var events []Event
	e := &event{}
	var data string

	for rows.Next() {
		if err = rows.Scan(&e.eventID, &e.kind, &e.version, &e.firedAt, &data); err != nil {
			return nil, err
		}

		events = append(events, &event{
			eventID:     e.eventID,
			data:        Data(data),
			firedAt:     e.firedAt,
			version:     e.version,
			kind:        e.kind,
			aggregateID: e.aggregateID,
		})
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return events, nil
}
