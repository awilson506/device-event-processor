-- +goose Up
-- +goose StatementBegin
CREATE VIEW latest_device_details AS 
	SELECT 
	(SELECT device_id
		FROM device_events
		WHERE device_id = b.device_name
		AND speed IS NOT NULL
		ORDER BY generated_at DESC
		LIMIT 1
	) AS device_id, 
	(SELECT generated_at
		FROM device_events
		WHERE device_id = b.device_name
		AND speed IS NOT NULL
		ORDER BY generated_at DESC
		LIMIT 1
	) AS generated_at, 
	(SELECT speed
		FROM device_events
		WHERE device_id = b.device_name
		AND speed IS NOT NULL
		ORDER BY generated_at DESC
		LIMIT 1
	) AS speed,
	(SELECT heading
		FROM device_events
		WHERE device_id = b.device_name
		AND heading IS NOT NULL
		ORDER BY generated_at DESC
		LIMIT 1
	) AS heading,
	(SELECT latitude
		FROM device_events d
		JOIN positions p
		ON d.id = p.device_id
		WHERE d.device_id = b.device_name
		AND latitude IS NOT NULL
		ORDER BY generated_at DESC
		LIMIT 1
	) AS latitude, 
	(SELECT longitude
		FROM device_events d
		JOIN positions p
		ON d.id = p.device_id
		WHERE d.device_id = b.device_name
		AND longitude IS NOT NULL
		ORDER BY generated_at DESC
		LIMIT 1
	) AS longitude
	FROM (SELECT DISTINCT d.device_id AS device_name
	FROM device_events d
	JOIN positions p
	ON d.id = p.device_id
	) AS b;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW latest_devices;
-- +goose StatementEnd
