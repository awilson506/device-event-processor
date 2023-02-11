package database

// GetLatestDeviceDetailsSQL - get the latest device details from the view
const GetLatestDeviceDetailsSQL = `
SELECT 
	device_id AS device,
	generated_at AS generated,
	speed,
	heading,
	latitude,
	longitude
FROM latest_device_details
WHERE device_id = $1;
`

const GetAllLatestDeviceDetailsSQL = `
SELECT 
	device_id AS device,
	generated_at AS generated,
	speed,
	heading,
	latitude,
	longitude
FROM latest_device_details
`

// InsertDeviceDetails - insert a device event, maybe with some position data
const InsertDeviceDetails = `
WITH new_device_details AS (
	INSERT INTO device_events (device_id, generated_at, speed, heading) 
	VALUES ($1, $2, $3, $4)
	RETURNING id
  )

INSERT INTO positions (device_id, latitude, longitude)
VALUES (
	(SELECT id FROM new_device_details), 
	$5, 
	$6
);
`
