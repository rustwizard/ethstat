-- +migrate Up
CREATE TABLE eth_txs (
    id bigserial,
    block_num bigint,
    tx_id text,
    from_addr text,
    to_addr text,
    value numeric(78,0),
    PRIMARY KEY (id)
);

CREATE INDEX eth_txs_block_num_tx_id_from_addr_index
    ON eth_txs (block_num, tx_id, from_addr);


-- +migrate Down
DROP TABLE eth_txs;