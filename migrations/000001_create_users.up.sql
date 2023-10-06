BEGIN;
CREATE TABLE IF NOT EXISTS users (
    id          BIGSERIAL   PRIMARY KEY,
    nama        VARCHAR NOT NULL,
    nik         VARCHAR NOT NULL,
    no_hp       VARCHAR NOT NULL,
    no_rekening VARCHAR NOT NULL,
    last_saldo  money
);

CREATE UNIQUE INDEX IF NOT EXISTS users_nik_no_hp_key
ON users(nik, no_hp);

CREATE UNIQUE INDEX IF NOT EXISTS users_no_rekening_key
ON users(no_rekening);
COMMIT;