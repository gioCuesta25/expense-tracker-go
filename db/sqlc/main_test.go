package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gioSmith25/expense-tracker/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")

	if err != nil {
		log.Fatal("Cannot load env: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
