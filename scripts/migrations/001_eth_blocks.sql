-- +migrate Up
CREATE TABLE eth_blocks (
    id bigserial,
    block_num bigint,
    txs text[],
    PRIMARY KEY (block_num)
);


-- +migrate Down
DROP TABLE eth_blocks;