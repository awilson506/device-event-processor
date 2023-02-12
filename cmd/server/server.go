package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/gocariq/golang_data_challenge/server"
	_ "github.com/lib/pq"
)

const DBCONN = "postgres://postgres:postgres@localhost:5432/process_db?sslmode=disable"

func main() {
	flag.Parse()

	db, dbErr := sql.Open("postgres", DBCONN)
	if dbErr != nil {
		panic(dbErr)
	}

	// start a new instance of the server
	s := server.New(db)
	err := s.Start()

	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
