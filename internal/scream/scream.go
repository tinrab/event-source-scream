package scream

import (
	"strings"
	"time"

	"github.com/tinrab/kit/id"
)

type Scream struct {
	ID        id.ID
	CreatedAt time.Time
	UserID    id.ID
	Body      string
}

func (s Scream) IsValid() bool {
	body := strings.TrimSpace(s.Body)

	if len(body) == 0 || len(body) > 160 {
		return false
	}

	if s.UserID == 0 {
		return false
	}

	return true
}
