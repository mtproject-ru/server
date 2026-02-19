-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_to_evens(
    user_id BIGINT NOT NULL,
    event_id BIGINT NOT NULL,

    PRIMARY KEY (user_id, event_id)
);

CREATE TABLE IF NOT EXISTS rooms_to_users(
    room_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,

    PRIMARY KEY (room_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_to_evens CASCADE;
DROP TABLE rooms_to_users CASCADE;
-- +goose StatementEnd
