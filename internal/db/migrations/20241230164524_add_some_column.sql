-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
  user_id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  age INT,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO "users" (email, password, first_name, last_name, age)
VALUES
  ('test1@example.com', 'password123', 'John', 'Doe', 30),
  ('test2@example.com', 'password123', 'Jane', 'Doe', 28),
  ('test3@example.com', 'password123', 'Alice', 'Smith', 25);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users" CASCADE;
-- +goose StatementEnd
