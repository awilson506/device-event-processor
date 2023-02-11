-- +goose Up
-- +goose StatementBegin
CREATE TABLE device_events (
    id SERIAL PRIMARY KEY,
    device_id VARCHAR(10) NOT NULL,
    generated_at TIMESTAMP NOT NULL,
    speed REAL DEFAULT NULL,
    heading INTEGER DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE device_events;
-- +goose StatementEnd
