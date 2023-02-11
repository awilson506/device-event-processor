package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	devices "github.com/gocariq/golang_data_challenge/pkg"
	_ "github.com/lib/pq"
)

const DBCONN = "postgres://postgres:postgres@localhost:5432/process_db?sslmode=disable"

type CustomTime struct {
	time.Time
}

func main() {
	// Open and ping a database connection to the docker db image.
	// Use this db connection if you like, it requires that the docker image
	// is running (via docker compose up).
	db, err := sql.Open("postgres", DBCONN)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	devicesService := devices.New(db)

	// TODO: add validation for args
	value := os.Args[1]

	record, err := devicesService.ProcessDeviceUpdate(value)
	if err != nil {
		panic(err)
	}

	jsonDeviceDetails, _ := json.Marshal(record)
	fmt.Printf("%s\n", jsonDeviceDetails)
}
