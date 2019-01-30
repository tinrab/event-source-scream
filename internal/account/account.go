package account

import (
	"time"

	"github.com/tinrab/kit/id"
)

type Account struct {
	ID        id.ID
	Title     string
	CreatedAt time.Time
}
