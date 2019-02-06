package user

import (
	"encoding/json"

	"github.com/tinrab/event-source-scream/internal/pkg/command"

	"github.com/tinrab/event-source-scream/internal/pkg/bus"
	userCommand "github.com/tinrab/event-source-scream/internal/user/command"
)

type Transport struct {
	bus     *bus.Bus
	service *Service
}

func NewTransport(b *bus.Bus, s *Service) *Transport {
	return &Transport{
		bus:     b,
		service: s,
	}
}

func (t *Transport) Run() error {
	if err := t.bus.HandleCommand(userCommand.KindCreateUser, t.handleCreateUser); err != nil {
		return err
	}
	return nil
}

func (t *Transport) handleCreateUser(_ string, data []byte) command.Result {
	var cmd userCommand.CreateUser
	if err := json.Unmarshal(data, &cmd); err != nil {
		return command.NewErrorResult(err)
	}

	u, err := t.service.CreateUser(cmd)
	if err != nil {
		return command.NewErrorResult(err)
	}

	return command.NewResult(userCommand.CreateUserResult{
		ID:   u.ID,
		Name: u.Name,
	})
}
