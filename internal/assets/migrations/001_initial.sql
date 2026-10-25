-- +migrate Up

-- +migrate StatementBegin
CREATE FUNCTION set_updated_at() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW() AT TIME ZONE 'utc';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TYPE recovery_event_type AS ENUM ('requested', 'cancelled', 'contested', 'finalized', 'locked', 'unlocked');
CREATE TYPE finalizer_submission_status AS ENUM
    ('submitting', 'submitted', 'confirmed', 'failed', 'rejected', 'unknown', 'resolved', 'expired');

-- Only finalized blocks are stored.
CREATE TABLE block
(
    block_id     SERIAL PRIMARY KEY,
    chain_id     INTEGER      NOT NULL,
    block_number BIGINT       NOT NULL CHECK (block_number >= 0),
    block_hash   TEXT         NOT NULL CHECK (block_hash ~ '^0x[0-9a-f]{64}$'),
    mined_at     TIMESTAMP(3) NOT NULL,
    is_active    BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at   TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at   TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);

CREATE UNIQUE INDEX block_chain_id_block_hash_idx ON block (chain_id, block_hash) WHERE is_active;
CREATE UNIQUE INDEX block_chain_id_block_number_idx ON block (chain_id, block_number) WHERE is_active;

CREATE TABLE transaction
(
    transaction_id SERIAL PRIMARY KEY,
    block_id       INTEGER      NOT NULL REFERENCES block (block_id),
    tx_hash        TEXT         NOT NULL CHECK (tx_hash ~ '^0x[0-9a-f]{64}$'),
    is_active      BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at     TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at     TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);

CREATE INDEX transaction_block_id_idx ON transaction (block_id);
CREATE UNIQUE INDEX transaction_block_id_tx_hash_idx ON transaction (block_id, tx_hash) WHERE is_active;

-- block_id is the indexing checkpoint, NULL before the first range.
CREATE TABLE contract
(
    contract_id SERIAL PRIMARY KEY,
    chain_id    INTEGER      NOT NULL,
    address     TEXT         NOT NULL CHECK (address ~ '^0x[0-9a-f]{40}$'),
    block_id    INTEGER REFERENCES block (block_id),
    is_active   BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at  TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
);

CREATE INDEX contract_block_id_idx ON contract (block_id);
CREATE UNIQUE INDEX contract_chain_id_address_idx ON contract (chain_id, address) WHERE is_active;

CREATE TABLE recovery_event
(
    recovery_event_id SERIAL PRIMARY KEY,
    contract_id       INTEGER             NOT NULL REFERENCES contract (contract_id),
    transaction_id    INTEGER             NOT NULL REFERENCES transaction (transaction_id),
    log_index         INTEGER             NOT NULL CHECK (log_index >= 0),
    event_type        recovery_event_type NOT NULL,
    passport_hash     TEXT                NOT NULL CHECK (passport_hash ~ '^0x[0-9a-f]{64}$'),
    old_identity_key  NUMERIC(78, 0) CHECK (old_identity_key > 0 AND old_identity_key < 2::NUMERIC ^ 256),
    new_identity_key  NUMERIC(78, 0) CHECK (new_identity_key > 0 AND new_identity_key < 2::NUMERIC ^ 256),
    dg_commit         NUMERIC(78, 0) CHECK (dg_commit >= 0 AND dg_commit < 2::NUMERIC ^ 256),
    ready_at          TIMESTAMP(3),
    locked_until      TIMESTAMP(3),
    is_active         BOOLEAN             NOT NULL DEFAULT TRUE,
    created_at        TIMESTAMP(3)        NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at        TIMESTAMP(3)        NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),

    CHECK ((new_identity_key IS NOT NULL) = (event_type = 'requested')),
    CHECK ((ready_at IS NOT NULL) = (event_type = 'requested')),
    CHECK ((old_identity_key IS NULL) = (dg_commit IS NULL)),
    CHECK (old_identity_key IS NULL OR event_type = 'requested'),
    CHECK (locked_until IS NULL OR event_type IN ('contested', 'cancelled')),
    CHECK (locked_until IS NOT NULL OR event_type <> 'contested')
);

CREATE INDEX recovery_event_contract_id_passport_hash_idx ON recovery_event (contract_id, passport_hash, event_type);
CREATE INDEX recovery_event_transaction_id_idx ON recovery_event (transaction_id);
CREATE UNIQUE INDEX recovery_event_transaction_id_log_index_idx ON recovery_event (transaction_id, log_index) WHERE is_active;

CREATE TABLE recovery_state
(
    recovery_state_id  SERIAL PRIMARY KEY,
    contract_id        INTEGER      NOT NULL REFERENCES contract (contract_id),
    recovery_event_id  INTEGER      NOT NULL REFERENCES recovery_event (recovery_event_id),
    passport_hash      TEXT         NOT NULL CHECK (passport_hash ~ '^0x[0-9a-f]{64}$'),
    old_identity_key   NUMERIC(78, 0) CHECK (old_identity_key > 0 AND old_identity_key < 2::NUMERIC ^ 256),
    new_identity_key   NUMERIC(78, 0) CHECK (new_identity_key > 0 AND new_identity_key < 2::NUMERIC ^ 256),
    dg_commit          NUMERIC(78, 0) CHECK (dg_commit >= 0 AND dg_commit < 2::NUMERIC ^ 256),
    ready_at           TIMESTAMP(3),
    locked_until       TIMESTAMP(3),
    is_recovery_locked BOOLEAN      NOT NULL DEFAULT FALSE,
    is_active          BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at         TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at         TIMESTAMP(3) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),

    CHECK ((new_identity_key IS NULL) = (ready_at IS NULL)),
    CHECK ((old_identity_key IS NULL) = (dg_commit IS NULL)),
    CHECK (old_identity_key IS NULL OR new_identity_key IS NOT NULL),
    -- A lock drops the pending request.
    CHECK (new_identity_key IS NULL OR NOT is_recovery_locked)
);

CREATE UNIQUE INDEX recovery_state_contract_id_passport_hash_idx ON recovery_state (contract_id, passport_hash) WHERE is_active;
CREATE INDEX recovery_state_recovery_event_id_idx ON recovery_state (recovery_event_id);
CREATE INDEX recovery_state_contract_id_ready_at_idx ON recovery_state (contract_id, ready_at);

CREATE TABLE finalizer_submission
(
    finalizer_submission_id SERIAL PRIMARY KEY,
    recovery_event_id       INTEGER                     NOT NULL REFERENCES recovery_event (recovery_event_id),
    status                  finalizer_submission_status NOT NULL,
    tx_hash                 TEXT CHECK (tx_hash ~ '^0x[0-9a-f]{64}$'),
    error                   TEXT                        NOT NULL DEFAULT '',
    is_active               BOOLEAN                     NOT NULL DEFAULT TRUE,
    created_at              TIMESTAMP(3)                NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    updated_at              TIMESTAMP(3)                NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),

    CHECK (tx_hash IS NOT NULL OR status NOT IN ('submitted', 'confirmed', 'failed'))
);

CREATE INDEX finalizer_submission_recovery_event_id_idx ON finalizer_submission (recovery_event_id);
CREATE INDEX finalizer_submission_status_idx ON finalizer_submission (status);
-- One blocking attempt per request.
CREATE UNIQUE INDEX finalizer_submission_blocking_idx ON finalizer_submission (recovery_event_id)
    WHERE is_active AND status NOT IN ('failed', 'rejected', 'expired');

CREATE TRIGGER block_set_updated_at BEFORE UPDATE ON block FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER transaction_set_updated_at BEFORE UPDATE ON transaction FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER contract_set_updated_at BEFORE UPDATE ON contract FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER recovery_event_set_updated_at BEFORE UPDATE ON recovery_event FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER recovery_state_set_updated_at BEFORE UPDATE ON recovery_state FOR EACH ROW EXECUTE FUNCTION set_updated_at();
CREATE TRIGGER finalizer_submission_set_updated_at BEFORE UPDATE ON finalizer_submission FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE VIEW recovery_state_detail AS
SELECT s.recovery_state_id,
       c.chain_id,
       c.address,
       s.passport_hash,
       s.old_identity_key,
       s.new_identity_key,
       s.dg_commit,
       s.ready_at,
       s.locked_until,
       s.is_recovery_locked,
       b.block_number AS updated_at_block,
       s.updated_at
FROM recovery_state s
         JOIN contract c ON c.contract_id = s.contract_id
         JOIN recovery_event e ON e.recovery_event_id = s.recovery_event_id
         JOIN transaction t ON t.transaction_id = e.transaction_id
         JOIN block b ON b.block_id = t.block_id
WHERE s.is_active;

-- +migrate Down

DROP VIEW recovery_state_detail;
DROP TABLE finalizer_submission;
DROP TABLE recovery_state;
DROP TABLE recovery_event;
DROP TABLE contract;
DROP TABLE transaction;
DROP TABLE block;
DROP TYPE finalizer_submission_status;
DROP TYPE recovery_event_type;
DROP FUNCTION set_updated_at();
