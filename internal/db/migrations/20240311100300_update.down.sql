-- reverse: set comment to column: "deleted_at" on table: "users"
COMMENT ON COLUMN "users" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "users"
COMMENT ON COLUMN "users" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "users"
COMMENT ON COLUMN "users" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "users"
COMMENT ON COLUMN "users" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "users"
COMMENT ON COLUMN "users" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "users"
COMMENT ON COLUMN "users" ."id" IS '';
-- reverse: create index "users_phone_key" to table: "users"
DROP INDEX "users_phone_key";
-- reverse: create "users" table
DROP TABLE "users";
