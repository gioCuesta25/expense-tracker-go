package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://user:secret@localhost:5433/expense_tracker?sslmode=disable"
	dbDriver = "postgres"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
