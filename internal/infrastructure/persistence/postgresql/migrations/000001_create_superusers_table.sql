-- +goose Up
CREATE TABLE IF NOT EXISTS superusers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username   TEXT NOT NULL,
    surname    TEXT NOT NULL,
    patronymic TEXT NOT NULL,
    password   TEXT NOT NULL,
    email      TEXT NOT NULL UNIQUE,
    role       TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_superusers_username ON superusers (username);

-- +goose Down
DROP INDEX IF EXISTS idx_superusers_username;
DROP TABLE IF EXISTS superusers;
