package main

import (
	"log"

	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-store/internal/pkg/bus"
	"github.com/tinrab/event-source-store/internal/pkg/config"
	"github.com/tinrab/event-source-store/internal/pkg/database"
	"github.com/tinrab/event-source-store/internal/user"
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

	r := user.NewRepository(db)
	idg := id.NewGenerator(0)
	s := user.NewService(idg, r)

	t := user.NewTransport(b, s)
	if err := t.Run(); err != nil {
		log.Fatal(err)
	}

	wait := make(chan struct{})
	<-wait
}
