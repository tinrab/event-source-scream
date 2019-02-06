package command

import (
	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/command"
)

const (
	KindCreateUser = "user.create"
)

type CreateUser struct {
	command.Command
	Name string `json:"name"`
}

type CreateUserResult struct {
	ID   id.ID  `json:"id"`
	Name string `json:"name"`
}
