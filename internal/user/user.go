package user

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/event"
	userCommand "github.com/tinrab/event-source-scream/internal/user/command"
	userEvent "github.com/tinrab/event-source-scream/internal/user/event"
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

func (u User) processCreate(cmd userCommand.CreateUser) ([]event.Event, error) {
	return []event.Event{
		event.New(userEvent.KindUserCreated, u.ID, userEvent.UserCreated{
			Name: cmd.Name,
		}),
	}, nil
}

func (u *User) apply(events []event.Event) {
	for _, e := range events {
		switch e.Kind {
		case userEvent.KindUserCreated:
			u.applyCreated(e)
		}
	}
}

func (u *User) applyCreated(e event.Event) {
	var data userEvent.UserCreated
	if err := json.Unmarshal(e.Data, &data); err != nil {
		log.Print(err)
		return
	}

	u.ID = e.AggregateID
	u.Name = data.Name
}
