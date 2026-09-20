-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS weather(
    chat_id int NOT NULL,
    city text NOT NULL,
    temp float NOT NULL,
    lon float  NOT NULL,
	lat float NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS weather;
-- +goose StatementEnd