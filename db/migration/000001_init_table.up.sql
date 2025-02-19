CREATE TABLE "users" (
  "user_id" UUID PRIMARY KEY NOT NULL UNIQUE,
  "frist_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "phone_number" VARCHAR UNIQUE NOT NULL,
  "password_hash" BYTEA NOT NULL,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "addresses" (
  "address_id" UUID PRIMARY KEY NOT NULL UNIQUE,
  "user_id" UUID NOT NULL,
  "address" VARCHAR(100) NOT NULL,
  "city" VARCHAR(50) NOT NULL,
  "state" VARCHAR(50) NOT NULL,
  "postal_code" VARCHAR(10) NOT NULL,
  "country" VARCHAR(50) NOT NULL,
  "is_default" BOOLEAN NOT NULL DEFAULT true,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "tokens" (
  "token_id" UUID PRIMARY KEY NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "token" UUID NOT NULL,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "expires_at" TIMESTAMP(0) NOT NULL
);

ALTER TABLE "addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "tokens" ADD FOREIGN KEY ("email") REFERENCES "users" ("email");