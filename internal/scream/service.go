package scream

import (
	"time"

	"github.com/pkg/errors"
	"github.com/tinrab/kit/id"
)

var (
	ErrInvalidScreamData = errors.New("invalid scream data")
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

func (s *Service) CreateScream(userID id.ID, body string) (*Scream, error) {
	sr := &Scream{
		ID:        s.idGenerator.Generate(),
		CreatedAt: time.Now().UTC(),
		UserID:    userID,
		Body:      body,
	}

	if !sr.IsValid() {
		return nil, ErrInvalidScreamData
	}

	if err := s.repository.SaveScream(*sr); err != nil {
		return sr, nil
	}

	return sr, nil
}

func (s *Service) ListScreamsByUser(userID id.ID) ([]Scream, error) {
	screams, err := s.repository.LoadScreamsByUser(userID)
	if err != nil {
		return nil, err
	}

	return screams, nil
}
