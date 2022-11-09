CREATE TABLE "emails_contract"
(
    "id"          bigserial PRIMARY KEY,
    "email_id"    bigserial   NOT NULL,
    "contract_id" bigserial   NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT group_unique UNIQUE (email_id, contract_id)
);

ALTER TABLE "emails_contract"
    ADD FOREIGN KEY ("email_id") REFERENCES "emails" ("id");

ALTER TABLE "emails_contract"
    ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id");
