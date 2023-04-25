CREATE TABLE "private_keys"
(
    "id"          bigserial PRIMARY KEY,
    "private_key" varchar     NOT NULL UNIQUE,
    "address"     varchar     NOT NULL UNIQUE,
    "updated_at"  timestamptz NOT NULL DEFAULT (now())
);
