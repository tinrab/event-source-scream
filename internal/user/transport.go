package user

import (
	"encoding/json"

	"github.com/tinrab/event-source-store/internal/api"
	"github.com/tinrab/event-source-store/internal/pkg/bus"
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
	if err := t.bus.Reply("api.user.create", t.createUser); err != nil {
		return err
	}
	return nil
}

func (t *Transport) createUser(data []byte) interface{} {
	var req api.CreateUserRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return api.CreateUserResponse{
			Error: err.Error(),
		}
	}

	u, err := t.service.CreateUser(req.Name)
	if err != nil {
		return api.CreateUserResponse{
			Error: err.Error(),
		}
	}

	return api.CreateUserResponse{
		ID:   uint64(u.ID),
		Name: u.Name,
	}
}
