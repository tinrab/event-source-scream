package user

import (
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

func (r *Repository) SaveUser(u User) error {
	_, err := r.db.Exec("INSERT INTO users(id, name) VALUES($1, $2)", u.ID, u.Name)
	return err
}

func (r *Repository) LoadUsers() ([]User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users ORDER BY id DESC")
	if err != nil {
		return nil, err
	}

	var users []User
	u := &User{}

	for rows.Next() {
		if err = rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, *u)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return users, nil
}
