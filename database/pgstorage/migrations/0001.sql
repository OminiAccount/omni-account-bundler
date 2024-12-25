-- +migrate Down
DROP SCHEMA IF EXISTS omni CASCADE;

-- +migrate Up
CREATE SCHEMA omni;

CREATE TABLE omni.user
(
    id          SERIAL8 PRIMARY KEY,
    owner       VARCHAR NOT NULL,
    account     VARCHAR,
    salt        INTEGER,
    failed_salt VARCHAR,
    create_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT user_unique UNIQUE (owner)
);

CREATE TABLE omni.user_info
(
    user_id    BIGINT REFERENCES omni.user (id) ON DELETE CASCADE,
    network_id INTEGER,
    nonce      BIGINT,
    gas        VARCHAR,
    create_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (user_id, network_id)
);

CREATE TABLE omni.batch
(
    batch_num           BIGINT PRIMARY KEY,
    batch_hash          VARCHAR,
    old_state_root      VARCHAR,
    state_root          VARCHAR,
    old_acc_input_hash  VARCHAR,
    acc_input_hash      VARCHAR,
    encoded             VARCHAR,
    status              INTEGER,
    create_at           TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at           TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE omni.proof
(
    batch_num       BIGINT NOT NULL REFERENCES omni.batch (batch_num) ON DELETE CASCADE,
    final_batch_num BIGINT NOT NULL REFERENCES omni.batch (batch_num) ON DELETE CASCADE,
    proof           VARCHAR,
    public_values   VARCHAR,
    create_at       TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at       TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE omni.operation
(
    id                         SERIAL8 PRIMARY KEY,
    user_id                    BIGINT REFERENCES omni.user (id) ON DELETE CASCADE,
    batch_num                  BIGINT,
    status                     INTEGER,
    signature                  VARCHAR,
    did                        VARCHAR,
    operation_type             INTEGER,
    operation_value            VARCHAR,
    phase                      INTEGER,
    exec_nonce                 BIGINT,
    exec_network_id            INTEGER,
    exec_calldata              BYTEA,
    exec_main_gas_limit        BIGINT,
    exec_main_gas_price        VARCHAR,
    exec_zkp_gas_limit         BIGINT,
    exec_dest_gas_limit        BIGINT,
    exec_dest_gas_price        VARCHAR,
    inner_exec_nonce           BIGINT,
    inner_exec_network_id      INTEGER,
    inner_exec_calldata        BYTEA,
    inner_exec_main_gas_limit  BIGINT,
    inner_exec_main_gas_price  VARCHAR,
    inner_exec_zkp_gas_limit   BIGINT,
    inner_exec_dest_gas_limit  BIGINT,
    inner_exec_dest_gas_price  VARCHAR,
    create_at                  TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at                  TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT sig_unique UNIQUE (signature)
);

CREATE TABLE omni.chain_height
(
    id         SERIAL8 PRIMARY KEY,
    network_id INTEGER,
    block_id   BIGINT,
    salt       BIGINT,
    create_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at  TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT net_unique UNIQUE (network_id)
);

CREATE TABLE omni.history (
    id            SERIAL8 PRIMARY KEY,
    did           VARCHAR,
    user_id       BIGINT REFERENCES omni.user (id) ON DELETE CASCADE,
    order_type    SMALLINT NOT NULL,
    from_info     jsonb,
    to_info       jsonb,
    source_chain  BIGINT NOT NULL,
    target_chain  BIGINT NOT NULL,
    source_hash   VARCHAR,
    target_hash   VARCHAR,
    swap_hash     VARCHAR,
    status        INTEGER,
    create_at     TIMESTAMP WITH TIME ZONE NOT NULL,
    update_at     TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE omni.hash_db (
    hash   VARCHAR PRIMARY KEY,
    data   jsonb
);

CREATE TABLE omni.hash_accval (
    hash   VARCHAR PRIMARY KEY,
    data   jsonb
);

CREATE TABLE omni.hash_keysource (
    hash   VARCHAR PRIMARY KEY,
    data   BYTEA
);

CREATE TABLE omni.hash_hashkey (
    hash   VARCHAR PRIMARY KEY,
    data   BYTEA
);

CREATE TABLE omni.hash_code (
    hash   VARCHAR PRIMARY KEY,
    data   BYTEA
);

CREATE TABLE omni.hash_info (
    lastroot VARCHAR,
    depth    BIGINT
);

-- Insert default values into hash_info table
INSERT INTO omni.hash_info (lastroot, depth) VALUES ('0x0', 0);