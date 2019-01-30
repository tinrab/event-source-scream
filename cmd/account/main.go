package main

import (
	"fmt"
	"log"

	"github.com/tinrab/event-source-store/internal/pkg/config"
	"github.com/tinrab/event-source-store/internal/pkg/database"
	"github.com/tinrab/event-source-store/internal/pkg/event"
)

func main() {
	c, err := config.Load("./cmd/account/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDatabase(c.Database)
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	eventStore := event.NewStore(db)

	//idGenerator := id.NewGenerator(0)
	//eventFactory := event.NewFactory(idGenerator)
	//err = eventStore.Save([]event.Event{
	//	eventFactory.Make("account.create", 42, account.Account{}),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	events, err := eventStore.Load(42, 3)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range events {
		fmt.Println(string(e.Data()))
	}
}
