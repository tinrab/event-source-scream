package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	"github.com/icrowley/fake"
	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-scream/internal/pkg/config"
	"github.com/tinrab/event-source-scream/internal/pkg/database"
	"github.com/tinrab/event-source-scream/internal/user"
)

func main() {
	flag.Parse()

	c, err := config.Load("./cmd/user/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDatabase(c.Database)
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	r := user.NewRepository(db)

	fake.Seed(42)
	rand.Seed(42)
	const n = 10

	for i := 0; i < n; i++ {
		u := user.User{
			ID:   id.ID(i + 1),
			Name: fmt.Sprintf("%s %s", fake.FirstName(), fake.LastName()),
		}

		if err := r.SaveUser(u); err != nil {
			log.Fatal(err)
		}
	}
}
