-- +goose Up
-- +goose StatementBegin
CREATE INDEX devices_id_generated_at ON device_events (id, generated_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX devices_id_generated_at;
-- +goose StatementEnd
