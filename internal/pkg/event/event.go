package event

import (
	"encoding/json"
	"time"

	"github.com/tinrab/kit/id"
)

type Event struct {
	ID          id.ID     `json:"id"`
	CommandID   id.ID     `json:"command_id"`
	AggregateID id.ID     `json:"aggregate_id"`
	Kind        string    `json:"kind"`
	Version     uint64    `json:"version"`
	FiredAt     time.Time `json:"fired_at"`
	Data        []byte    `json:"data"`
}

func New(kind string, aggregateID id.ID, data interface{}) Event {
	bd, _ := json.Marshal(data)

	return Event{
		AggregateID: aggregateID,
		Kind:        kind,
		Version:     1,
		FiredAt:     time.Now().UTC(),
		Data:        bd,
	}
}

func NewFromPrevious(previous Event, kind string, data interface{}) Event {
	bd, _ := json.Marshal(data)

	return Event{
		AggregateID: previous.AggregateID,
		Kind:        kind,
		Version:     previous.Version + 1,
		FiredAt:     time.Now().UTC(),
		Data:        bd,
	}
}
