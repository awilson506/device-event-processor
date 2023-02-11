-- +goose Up
-- +goose StatementBegin
CREATE TABLE positions (
    id SERIAL PRIMARY KEY,
    device_id INT,
    latitude DOUBLE PRECISION DEFAULT NULL,
    longitude DOUBLE PRECISION DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_device
        FOREIGN KEY(device_id)
            REFERENCES device_events(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE positions;
-- +goose StatementEnd
