CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "description" varchar(200) NOT NULL,
  "amount" bigserial NOT NULL,
  "category_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "incomes" (
  "id" bigserial PRIMARY KEY,
  "description" varchar(200) NOT NULL,
  "amount" bigserial NOT NULL,
  "category_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "is_for_incomes" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "expenses" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "incomes" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");