ALTER TABLE "addresses" DROP CONSTRAINT IF EXISTS addresses_user_id_fkey;

DROP TABLE IF EXISTS "addresses";
DROP TABLE IF EXISTS "session";
DROP TABLE IF EXISTS "users";