package account

import (
	"github.com/tinrab/event-source-store/internal/pkg/database"
)

type Repository interface {
	SaveAccount(a Account) error
	LoadAccounts() ([]Account, error)
}

type repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) SaveAccount(a Account) error {
	_, err := r.db.Exec("INSERT INTO accounts(id, title, created_at) VALUES($1, $2, $3)", a.ID, a.Title, a.CreatedAt)
	return err
}

func (r *repository) LoadAccounts() ([]Account, error) {
	rows, err := r.db.Query("SELECT id, title, created_at FROM accounts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	var accounts []Account
	a := &Account{}

	for rows.Next() {
		if err = rows.Scan(&a.ID, &a.Title, &a.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, *a)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return accounts, nil
}
