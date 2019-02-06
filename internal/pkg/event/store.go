package event

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/database"
)

type Store struct {
	db    *database.Database
	table string
}

func NewStore(db *database.Database, table string) *Store {
	return &Store{
		db:    db,
		table: table,
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

	stmt, err := txn.Prepare(pq.CopyIn(s.table, "id", "aggregate_id", "kind", "version", "fired_at", "data"))
	if err != nil {
		return
	}

	for _, e := range events {
		_, err = stmt.Exec(e.ID, e.AggregateID, e.Kind, e.Version, e.FiredAt, string(e.Data))
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
		fmt.Sprintf("SELECT id, kind, version, fired_at, data FROM %s WHERE aggregate_id = $1 AND version > $2 ORDER BY fired_at DESC", s.table),
		aggregateID,
		fromVersion,
	)
	if err != nil {
		return nil, err
	}

	var events []Event
	e := &Event{}
	var data string

	for rows.Next() {
		if err = rows.Scan(&e.ID, &e.Kind, &e.Version, &e.FiredAt, &data); err != nil {
			return nil, err
		}

		events = append(events, Event{
			ID:          e.ID,
			Data:        []byte(data),
			FiredAt:     e.FiredAt,
			Version:     e.Version,
			Kind:        e.Kind,
			AggregateID: e.AggregateID,
		})
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return events, nil
}
