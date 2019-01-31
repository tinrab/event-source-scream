package event

import (
	"encoding/json"
	"time"

	"github.com/tinrab/kit/id"
)

type Factory struct {
	idGenerator *id.Generator
}

func NewFactory(idGenerator *id.Generator) *Factory {
	return &Factory{
		idGenerator: idGenerator,
	}
}

func (f *Factory) Make(kind Kind, aggregateID id.ID, data interface{}) Event {
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

func (f *Factory) MakeFromPrevious(previous Event, kind Kind, data interface{}) Event {
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
