ALTER TABLE "addresses" DROP CONSTRAINT IF EXISTS addresses_user_id_fkey;
ALTER TABLE "tokens" DROP CONSTRAINT IF EXISTS tokens_user_id_fkey;

DROP TABLE IF EXISTS "addresses";
DROP TABLE IF EXISTS "tokens";
DROP TABLE IF EXISTS "users";