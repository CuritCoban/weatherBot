-- +goose Up

CREATE TABLE IF NOT EXISTS weather(
    chat_id int NOT NULL,
    city text NOT NULL,
    temp float NOT NULL,
    lon float  NOT NULL,
	lat float NOT NULL,
    created_at timestamp NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS weather;