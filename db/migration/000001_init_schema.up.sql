CREATE TABLE "accounts"(
                           "id" bigserial PRIMARY KEY,
                           "owner" varchar not null,
                           "balance" bigint not null,
                           "created_at" timestamp not null default (now())
);

CREATE TABLE "entries"(
                          "id" bigserial primary key,
                          "account_id" bigint not null,
                          "amount" bigint not null,
                          "created_at" timestamp not null default (now())
);

CREATE TABLE "transfers"(
                            "id" bigserial primary key,
                            "from_account_id" bigint not null,
                            "to_account_id" bigint not null,
                            "amount" bigint not null,
                            "created_at" timestamp not null default (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");
ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("owner");
CREATE INDEX ON "entries" ("account_id");
CREATE INDEX ON "transfers" ("from_account_id");
CREATE INDEX ON "transfers" ("to_account_id");
CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");
