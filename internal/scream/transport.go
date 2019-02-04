package scream

import (
	"encoding/json"

	"github.com/tinrab/event-source-scream/internal/api"
	"github.com/tinrab/event-source-scream/internal/pkg/bus"
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
	if err := t.bus.Reply("scream.create", t.createScream); err != nil {
		return err
	}
	if err := t.bus.Reply("scream.list", t.listScreams); err != nil {
		return err
	}
	return nil
}

func (t *Transport) createScream(data []byte) interface{} {
	var req api.CreateScreamRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return api.CreateScreamResponse{
			Error: err.Error(),
		}
	}

	s, err := t.service.CreateScream(req.UserID, req.Body)
	if err != nil {
		return api.CreateScreamResponse{
			Error: err.Error(),
		}
	}

	createdAt := s.CreatedAt

	return api.CreateScreamResponse{
		ID:        s.ID,
		CreatedAt: &createdAt,
		UserID:    s.UserID,
		Body:      s.Body,
	}
}

func (t *Transport) listScreams(data []byte) interface{} {
	var req api.ListScreamsRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return api.ListScreamsResponse{
			Error: err.Error(),
		}
	}

	screams, err := t.service.ListScreamsByUser(req.UserID)
	if err != nil {
		return api.ListScreamsResponse{
			Error: err.Error(),
		}
	}

	var sr []api.ScreamResponse
	for _, s := range screams {
		createdAt := s.CreatedAt
		sr = append(sr, api.ScreamResponse{
			ID:        s.ID,
			CreatedAt: &createdAt,
			Body:      s.Body,
		})
	}

	return api.ListScreamsResponse{
		UserID:  req.UserID,
		Screams: sr,
	}
}
