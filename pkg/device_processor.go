package devices

import (
	"database/sql"
	"encoding/json"
	"log"

	database "github.com/gocariq/golang_data_challenge/database"
)

// DeviceDB - only holds the db connection for now
type DeviceDB struct {
	Connection *sql.DB
}

// New - grab a new instance of the device db service
func New(db *sql.DB) *DeviceDB {
	return &DeviceDB{
		Connection: db,
	}
}

// ProcessDeviceUpdate - insert the new event and return the latest data
func (d *DeviceDB) ProcessDeviceUpdate(deviceJsonUpdate string) (*database.DeviceDetails, error) {
	newDeviceDetails := database.DeviceDetails{}
	err := json.Unmarshal([]byte(deviceJsonUpdate), &newDeviceDetails)

	if err != nil {
		log.Fatalf("invalid json: %s", err)
	}

	// in the case that a late record comes in and we want the view to pick up the non null value
	// we should store the event even if it's old to keep the record
	d.InsertDeviceDetails(&newDeviceDetails)

	return d.GetLatestDeviceDetails(newDeviceDetails.DeviceId), nil
}
