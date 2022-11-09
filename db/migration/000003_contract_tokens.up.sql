CREATE TABLE "tokens_contract"
(
    "id"          bigserial PRIMARY KEY,
    "contract_id" bigserial   NOT NULL,
    "name"        varchar     NOT NULL,
    "symbol"      varchar     NOT NULL,
    "address"     varchar(42) NOT NULL,
    "decimals"    bigint      NOT NULL,
    "updated_at"  timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT token_contract_unique UNIQUE (address, contract_id)
);

ALTER TABLE "tokens_contract"
    ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");
