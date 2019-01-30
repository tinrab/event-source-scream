package event

import (
	"time"

	"github.com/tinrab/kit/id"
)

type Kind string

type Data []byte

type Event interface {
	ID() id.ID
	AggregateID() id.ID
	Kind() Kind
	Version() uint64
	FiredAt() time.Time
	Data() Data
}

type event struct {
	eventID     id.ID
	aggregateID id.ID
	kind        Kind
	version     uint64
	firedAt     time.Time
	data        Data
}

func (e *event) ID() id.ID {
	return e.eventID
}

func (e *event) AggregateID() id.ID {
	return e.aggregateID
}

func (e *event) Kind() Kind {
	return e.kind
}

func (e *event) Version() uint64 {
	return e.version
}

func (e *event) FiredAt() time.Time {
	return e.firedAt
}

func (e *event) Data() Data {
	return e.data
}
