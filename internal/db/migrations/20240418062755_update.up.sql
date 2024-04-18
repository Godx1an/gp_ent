-- modify "admins" table
ALTER TABLE "admins" ADD COLUMN "email" character varying NOT NULL DEFAULT '';
-- modify "users" table
ALTER TABLE "users" ADD COLUMN "email" character varying NOT NULL DEFAULT '';
