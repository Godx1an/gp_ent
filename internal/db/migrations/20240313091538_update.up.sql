-- modify "users" table
ALTER TABLE "users" ADD COLUMN "school" character varying NOT NULL DEFAULT '', ADD COLUMN "next_update_time" timestamptz NULL;
