package command

import (
	"encoding/json"

	"github.com/tinrab/kit/id"
)

type Command struct {
	ID   id.ID  `json:"id"`
	Kind string `json:"kind"`
	Data []byte `json:"data"`
}

func New(id id.ID, kind string, data interface{}) Command {
	bd, _ := json.Marshal(data)

	return Command{
		ID:   id,
		Kind: kind,
		Data: bd,
	}
}
