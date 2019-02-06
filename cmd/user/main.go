package main

import (
	"log"

	"github.com/tinrab/event-source-scream/internal/pkg/event"
	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/bus"
	"github.com/tinrab/event-source-scream/internal/pkg/config"
	"github.com/tinrab/event-source-scream/internal/pkg/database"
	"github.com/tinrab/event-source-scream/internal/user"
)

func main() {
	c, err := config.Load("./cmd/user/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDatabase(c.Database)
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	b := bus.NewBus(c.Bus)
	if err := b.Open(); err != nil {
		log.Fatal(err)
	}

	es := event.NewStore(db, "events")

	r := user.NewRepository(db, es)
	idg := id.NewGenerator(1)
	s := user.NewService(idg, r)

	t := user.NewTransport(b, s)
	if err := t.Run(); err != nil {
		log.Fatal(err)
	}

	wait := make(chan struct{})
	<-wait
}
