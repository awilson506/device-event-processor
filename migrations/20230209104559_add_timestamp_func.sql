-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON device_events
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON positions
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER set_timestamp;
-- +goose StatementEnd
