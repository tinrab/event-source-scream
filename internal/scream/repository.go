package scream

import (
	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/database"
)

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) SaveScream(s Scream) error {
	_, err := r.db.Exec(
		"INSERT INTO screams(id, created_at, user_id, body) VALUES($1, $2, $3, $4)",
		s.ID,
		s.CreatedAt,
		s.UserID,
		s.Body,
	)
	return err
}

func (r *Repository) LoadScreamsByUser(userID id.ID) ([]Scream, error) {
	rows, err := r.db.Query("SELECT id, created_at, user_id, body FROM screams WHERE user_id = $1 ORDER BY id DESC", userID)
	if err != nil {
		return nil, err
	}

	var screams []Scream
	s := &Scream{}

	for rows.Next() {
		if err = rows.Scan(&s.ID, &s.CreatedAt, &s.UserID, &s.Body); err != nil {
			return nil, err
		}
		screams = append(screams, *s)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return screams, nil
}
