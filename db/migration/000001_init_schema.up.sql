CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "entries" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "description" varchar(200) NOT NULL,
  "amount" bigserial NOT NULL,
  "category_id" UUID NOT NULL,
  "type_id" UUID NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "is_for_incomes" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entry_types" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "entries" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("type_id") REFERENCES "entry_types" ("id");