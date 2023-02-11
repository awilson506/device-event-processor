package devices

import (
	"database/sql"
	"log"

	"github.com/gocariq/golang_data_challenge/database"
)

// GetLatestDeviceDetails - get the latest device from the view
func (d *DeviceDB) GetLatestDeviceDetails(deviceId string) *database.DeviceDetails {
	deviceDetails := database.DeviceDetails{}
	d.Connection.QueryRow(database.GetLatestDeviceDetailsSQL, deviceId).Scan(
		&deviceDetails.DeviceId,
		&deviceDetails.Generated,
		&deviceDetails.Speed,
		&deviceDetails.Heading,
		&deviceDetails.Position.Latitude,
		&deviceDetails.Position.Longitude,
	)
	// if we don't find a record just return the empty instance for the processor to use
	return &deviceDetails
}

// InsertDeviceDetails - insert a device event
func (d *DeviceDB) InsertDeviceDetails(device *database.DeviceDetails) *sql.Result {
	insertStatment, err := d.Connection.Prepare(database.InsertDeviceDetails)
	if err != nil {
		log.Fatalf("insert prepare failed for device details: %s", err)
	}

	result, err := insertStatment.Exec(
		device.DeviceId,
		device.Generated,
		device.Speed,
		device.Heading,
		device.Position.Latitude,
		device.Position.Longitude,
	)

	if err != nil {
		log.Fatalf("insert failed for device details: %s", err)
	}

	return &result
}
