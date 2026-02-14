DROP TABLE IF EXISTS transactions;

DROP TYPE IF EXISTS transaction_type_enum;

ALTER TABLE
  users DROP COLUMN IF EXISTS balance;