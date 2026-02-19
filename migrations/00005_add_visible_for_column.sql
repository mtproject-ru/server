-- +goose Up
-- +goose StatementBegin
ALTER TABLE events ADD COLUMN visible_for BIGINT ARRAY;
ALTER TABLE rooms ADD COLUMN visible_for BIGINT ARRAY;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE events DROP COLUMN visible_for;
ALTER TABLE rooms DROP COLUMN visible_for;
-- +goose StatementEnd
