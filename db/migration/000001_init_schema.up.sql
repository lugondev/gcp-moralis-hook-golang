CREATE TABLE "emails"
(
    "id"    bigserial PRIMARY KEY,
    "name"  varchar NOT NULL,
    "email" varchar NOT NULL UNIQUE CHECK (email ~* '^.+@.+\..+$'
) ,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "contracts"
(
    "id"          bigserial PRIMARY KEY,
    "name"        varchar     NOT NULL,
    "is_contract" bool        NOT NULL DEFAULT true,
    "chain_id"    numeric     NOT NULL,
    "address"     varchar(42) NOT NULL UNIQUE CHECK (address ~* '^0x[a-fA-F0-9]{40}$'
) ,
    "network"    varchar     NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
