package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/icrowley/fake"
	"github.com/tinrab/kit/id"

	"github.com/tinrab/event-source-store/internal/account"
	"github.com/tinrab/event-source-store/internal/pkg/config"
	"github.com/tinrab/event-source-store/internal/pkg/database"
)

func main() {
	flag.Parse()

	c, err := config.Load("./cmd/account/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewDatabase(c.Database)
	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	r := account.NewRepository(db)

	const n = 10
	idGenerator := id.NewGenerator(0)
	ids := idGenerator.GenerateList(n)

	for i := 0; i < n; i++ {
		a := account.Account{
			ID:        ids[i],
			Title:     fmt.Sprintf("%s %s", fake.FirstName(), fake.LastName()),
			CreatedAt: time.Now().UTC().Add(-time.Duration(rand.Int31n(100)) * time.Hour),
		}

		if err := r.SaveAccount(a); err != nil {
			log.Fatal(err)
		}
	}
}
