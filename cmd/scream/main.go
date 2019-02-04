package main

import (
	"log"

	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/scream"

	"github.com/tinrab/event-source-scream/internal/pkg/bus"
	"github.com/tinrab/event-source-scream/internal/pkg/config"
	"github.com/tinrab/event-source-scream/internal/pkg/database"
)

func main() {
	c, err := config.Load("./cmd/scream/config.yml")
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

	r := scream.NewRepository(db)
	idg := id.NewGenerator(1)
	s := scream.NewService(idg, r)

	t := scream.NewTransport(b, s)
	if err := t.Run(); err != nil {
		log.Fatal(err)
	}

	wait := make(chan struct{})
	<-wait
}
