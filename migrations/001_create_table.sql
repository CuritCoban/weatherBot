-- goose+ Up

CREATE TABLE IF NOT EXISTS weather(
    id serial PRIMARY KEY,
    city text NOT NULL,
    temp float NOT NULL
)
-- goose+ Down

DROP TABLE IF EXISTS weather;