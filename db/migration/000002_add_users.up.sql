CREATE TABLE IF NOT EXISTS "users" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "hashed_password" varchar NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "full_name" varchar(50) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE entries ADD user_id UUID NOT NULL;

ALTER TABLE "entries" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");