-- +goose Up
-- +goose StatementBegin
-- USERS
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


-- WALLETS
CREATE TYPE currency_type AS ENUM ('USD', 'UAH', 'EUR');

CREATE TABLE IF NOT EXISTS "wallets" (
	wallet_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	balance INT NOT NULL DEFAULT 0,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	currency currency_type DEFAULT 'USD',
	FOREIGN KEY (user_id) REFERENCES "users" (user_id)
);

INSERT INTO "wallets" (user_id, balance, name, description, currency)
VALUES
	(1, 1000, 'Main Wallet', 'My main wallet', 'USD'),
	(2, 500, 'Savings Wallet', 'My savings wallet', 'UAH'),
	(3, 2000, 'Travel Wallet', 'My travel wallet', 'EUR');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "wallets" CASCADE;
DROP TYPE IF EXISTS currency_type;
-- +goose StatementEnd
