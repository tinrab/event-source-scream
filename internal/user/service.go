package user

import (
	"github.com/pkg/errors"
	"github.com/tinrab/kit/id"
)

var (
	ErrInvalidUserData = errors.New("invalid user data")
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

func (s *Service) CreateUser(name string) (*User, error) {
	u := &User{
		ID:   s.idGenerator.Generate(),
		Name: name,
	}

	if !u.IsValid() {
		return nil, ErrInvalidUserData
	}

	if err := s.repository.SaveUser(*u); err != nil {
		return nil, err
	}

	return u, nil
}
