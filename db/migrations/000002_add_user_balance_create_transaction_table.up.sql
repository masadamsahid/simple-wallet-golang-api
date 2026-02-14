ALTER TABLE
  users
ADD
  COLUMN balance BIGINT NOT NULL DEFAULT 0;

CREATE TYPE transaction_type_enum AS ENUM ('withdrawal', 'deposit');

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  amount BIGINT NOT NULL,
  transaction_type transaction_type_enum NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NULL
);

ALTER TABLE
  transactions
ADD
  CONSTRAINT fk_transaction_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT;

CREATE INDEX idx_transactions_user_id ON transactions(user_id);