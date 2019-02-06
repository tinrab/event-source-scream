package command

import (
	"encoding/json"
)

type Result struct {
	Error string `json:"error"`
	Data  []byte `json:"data"`
}

func NewResult(data interface{}) Result {
	bd, _ := json.Marshal(data)

	return Result{
		Data: bd,
	}
}

func NewErrorResult(err error) Result {
	return Result{
		Error: err.Error(),
	}
}

func (r Result) IsError() bool {
	return r.Error != ""
}
