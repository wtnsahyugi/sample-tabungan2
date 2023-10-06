BEGIN;
CREATE TABLE IF NOT EXISTS transactions (
    id                  BIGSERIAL   PRIMARY KEY,
    transaction_code    VARCHAR NOT NULL,
    no_rekening         VARCHAR NOT NULL,
    transaction_type    VARCHAR NOT NULL,
    created_date        TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS transactions_no_rekening_idx ON transactions(no_rekening);
COMMIT;