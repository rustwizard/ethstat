-- +migrate Up
CREATE TABLE eth_txs (
    id bigserial,
    block_num bigint,
    tx_id text,
    from_addr text,
    to_addr text,
    value numeric(78,0),
    PRIMARY KEY (block_num, tx_id)
);


-- +migrate Down
DROP TABLE eth_txs;