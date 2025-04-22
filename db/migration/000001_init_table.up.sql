CREATE TABLE "users" (
  "user_id" UUID PRIMARY KEY NOT NULL UNIQUE,
  "frist_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "email" VARCHAR UNIQUE NOT NULL,
  "phone_number" VARCHAR NOT NULL,
  "password_hash" BYTEA NOT NULL,
  "role" VARCHAR NOT NULL,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "addresses" (
  "address_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "flat_floor" VARCHAR(50),
  "building" VARCHAR(255),
  "street" VARCHAR(255) NOT NULL,
  "district" VARCHAR(100) NOT NULL,
  "region" VARCHAR(50) NOT NULL,
  "country" VARCHAR(50) NOT NULL,
  "zip_code" VARCHAR(50),
  "is_default" BOOLEAN NOT NULL DEFAULT false,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "session" (
  "session_id" UUID PRIMARY KEY NOT NULL,
  "email" VARCHAR NOT NULL,
  "token" VARCHAR NOT NULL,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "expires_at" TIMESTAMP(0) NOT NULL
);

ALTER TABLE "addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
