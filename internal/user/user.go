package user

import (
	"strings"

	"github.com/tinrab/kit/id"
)

type User struct {
	ID   id.ID
	Name string
}

func (u User) IsValid() bool {
	name := strings.TrimSpace(u.Name)

	if len(name) < 3 || len(name) > 32 {
		return false
	}

	return true
}
