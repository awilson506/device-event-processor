package devices

import (
	"database/sql"
	"encoding/json"
	"log"

	database "github.com/gocariq/golang_data_challenge/database"
)

// DeviceDB - only holds the db connection for now
type DeviceService struct {
	Connection *sql.DB
}

// New - grab a new instance of the device db service
func New(db *sql.DB) *DeviceService {
	return &DeviceService{
		Connection: db,
	}
}

// ProcessDeviceUpdate - insert the new event and return the latest data
func (d *DeviceService) ProcessDeviceUpdate(deviceJsonUpdate string) (*database.DeviceDetails, error) {
	newDeviceDetails := database.DeviceDetails{}
	err := json.Unmarshal([]byte(deviceJsonUpdate), &newDeviceDetails)

	if err != nil {
		log.Fatalf("invalid json: %s", err)
	}

	// in the case that a late record comes in and we want the view to pick up the non null value
	// most likely we would be processing these off a queue which doesn't always ensure order...
	d.InsertDeviceDetails(&newDeviceDetails)

	return d.GetLatestDeviceDetails(newDeviceDetails.DeviceId), nil
}
