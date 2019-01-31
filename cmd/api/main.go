package main

import (
	"log"

	"github.com/tinrab/event-source-store/internal/api"
	"github.com/tinrab/event-source-store/internal/pkg/bus"
	"github.com/tinrab/event-source-store/internal/pkg/config"
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

	s := api.NewServer(c.API.Port, b)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}