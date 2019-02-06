package user

import (
	"github.com/tinrab/kit/id"

	userCommand "github.com/tinrab/event-source-scream/internal/user/command"
)

type Service struct {
	idGenerator *id.Generator
	repository  *Repository
}

func NewService(idg *id.Generator, r *Repository) *Service {
	return &Service{
		idGenerator: idg,
		repository:  r,
	}
}

func (s *Service) CreateUser(cmd userCommand.CreateUser) (*User, error) {
	u := &User{
		ID: s.idGenerator.Generate(),
	}

	events, err := u.processCreate(cmd)
	if err != nil {
		return nil, err
	}

	u.apply(events)

	return u, nil
}
