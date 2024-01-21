package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://user:secret@localhost:5433/expense_tracker?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db")
	}
	/*
		La función New crea una instancia de Queries
		para poder interactuar con la base de datos
	*/
	testQueries = New(conn)

	os.Exit(m.Run())
}
