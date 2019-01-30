package event

import (
	"encoding/json"
	"time"

	"github.com/tinrab/kit/id"
)

type Factory interface {
	Make(kind Kind, aggregateID id.ID, data interface{}) Event
}

type factory struct {
	idGenerator *id.Generator
}

func NewFactory(idGenerator *id.Generator) Factory {
	return &factory{
		idGenerator: idGenerator,
	}
}

func (f *factory) Make(kind Kind, aggregateID id.ID, data interface{}) Event {
	bd, _ := json.Marshal(data)

	return &event{
		eventID:     f.idGenerator.Generate(),
		aggregateID: aggregateID,
		kind:        kind,
		version:     1,
		firedAt:     time.Now().UTC(),
		data:        bd,
	}
}

func (f *factory) MakeFromPrevious(previous Event, kind Kind, data interface{}) Event {
	bd, _ := json.Marshal(data)

	return &event{
		eventID:     f.idGenerator.Generate(),
		aggregateID: previous.AggregateID(),
		kind:        kind,
		version:     previous.Version() + 1,
		firedAt:     time.Now().UTC(),
		data:        bd,
	}
}
