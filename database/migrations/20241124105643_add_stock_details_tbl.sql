-- +goose Up
-- Create stock_details table
CREATE TABLE IF NOT EXISTS stock_details (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- +goose Down
-- Drop stock_details table
DROP TABLE IF EXISTS stock_details;