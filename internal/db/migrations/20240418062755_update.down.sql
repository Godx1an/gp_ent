-- reverse: modify "users" table
ALTER TABLE "users" DROP COLUMN "email";
-- reverse: modify "admins" table
ALTER TABLE "admins" DROP COLUMN "email";
