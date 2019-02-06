package main

import (
	"log"

	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/api"
	"github.com/tinrab/event-source-scream/internal/pkg/bus"
	"github.com/tinrab/event-source-scream/internal/pkg/config"
)

func main() {
	c, err := config.Load("./cmd/api/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	b := bus.NewBus(c.Bus)
	if err := b.Open(); err != nil {
		log.Fatal(err)
	}

	idg := id.NewGenerator(0)

	s := api.NewServer(c.API.Port, b, idg)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
