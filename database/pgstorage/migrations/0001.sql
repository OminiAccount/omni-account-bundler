-- +migrate Down
DROP SCHEMA IF EXISTS omni CASCADE;

-- +migrate Up
CREATE SCHEMA omni;

CREATE TABLE omni.user
(
    id        SERIAL8 PRIMARY KEY,
    owner     VARCHAR NOT NULL,
    account   VARCHAR,
    salt      INTEGER,
    create_at TIMESTAMP WITH TIME ZONE NOT NULL,
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

CREATE TABLE omni.operation
(
    id                         SERIAL8 PRIMARY KEY,
    user_id                    BIGINT REFERENCES omni.user (id) ON DELETE CASCADE,
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
    update_at                  TIMESTAMP WITH TIME ZONE NOT NULL
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
