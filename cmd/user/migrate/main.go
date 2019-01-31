package main

import (
	"flag"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/tinrab/event-source-store/internal/pkg/config"
	"github.com/tinrab/event-source-store/internal/pkg/database"
)

var action string

func init() {
	flag.StringVar(&action, "action", "up", "action to perform")
}

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

	driver, err := postgres.WithInstance(db.Connection(), &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/user/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	switch action {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "drop":
		if err := m.Drop(); err != nil {
			log.Fatal(err)
		}
	}
}
