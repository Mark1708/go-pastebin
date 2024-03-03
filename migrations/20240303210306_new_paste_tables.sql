-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS paste
(
    hash VARCHAR(6) NOT NULL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    author VARCHAR(100) NOT NULL,
    content_path VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    expiration VARCHAR(5) NOT NULL,
    CONSTRAINT hash_unique UNIQUE (hash)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS paste;
-- +goose StatementEnd
