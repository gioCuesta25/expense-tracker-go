package main

import (
	"database/sql"
	"log"

	"github.com/gioSmith25/expense-tracker/api"
	db "github.com/gioSmith25/expense-tracker/db/sqlc"
	"github.com/gioSmith25/expense-tracker/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load env config: ", err.Error())
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db")
	}

	q := db.New(conn)
	server := api.NewServer(q)

	err = server.Start(config.Host)

	if err != nil {
		log.Fatal("cannot possible init the server")
	}
}
