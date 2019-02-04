package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/tinrab/event-source-scream/internal/pkg/config"
)

type Database struct {
	cfg  config.DatabaseConfig
	conn *sql.DB
}

func NewDatabase(c config.DatabaseConfig) *Database {
	return &Database{
		cfg: c,
	}
}

func (db *Database) Open() error {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?port=%d&sslmode=disable",
		db.cfg.User,
		db.cfg.Password,
		db.cfg.Host,
		db.cfg.Name,
		db.cfg.Port,
	)
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	if err = conn.Ping(); err != nil {
		return err
	}

	db.conn = conn

	return nil
}

func (db *Database) Close() error {
	if db.conn == nil {
		return nil
	}
	if err := db.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *Database) Exec(statement string, args ...interface{}) (sql.Result, error) {
	return db.conn.Exec(statement, args...)
}

func (db *Database) Begin() (*sql.Tx, error) {
	return db.conn.Begin()
}

func (db *Database) Connection() *sql.DB {
	return db.conn
}
