# GoLang Device Event Aggregator 

A small tool to input device events and get back the latest data about the device

## Getting the service setup & running

Note this application requires Go version `1.19` but supports back to `1.18` if you already have a previous version
installed.  You will have to update the Go module to require a later version if you choose not to update.  This can be done by 
running: 
```sh
go mod edit -go=1.MY_VERSION
```

Start the Postgres database service: `docker compose`
```sh
docker compose up
```

Run the database migrations:
```sh
./bin/goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/process_db?sslmode=disable" up
```

Build the service:
```sh
go build process.go
```

Process the records:
```sh
./process '{"device": "A123", "generated": "2022-01-01 15:00:00.000", "speed": 48.7, "heading": 101}'
```
Expected output:
```json
{"device":"A123","heading":101,"speed":48.7}
```
