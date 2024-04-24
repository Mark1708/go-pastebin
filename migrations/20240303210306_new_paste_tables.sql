-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS paste
(
    hash       VARCHAR(8)  NOT NULL PRIMARY KEY,
    title      VARCHAR(50) NOT NULL,
    visibility VARCHAR(10),
    created_at TIMESTAMP   NOT NULL,
    expired_at TIMESTAMP   NOT NULL,
    content    TEXT        NOT NULL,
    CONSTRAINT hash_unique UNIQUE (hash)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS paste;
-- +goose StatementEnd
