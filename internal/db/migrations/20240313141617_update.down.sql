-- reverse: set comment to column: "deleted_at" on table: "admins"
COMMENT ON COLUMN "admins" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "admins"
COMMENT ON COLUMN "admins" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "admins"
COMMENT ON COLUMN "admins" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "admins"
COMMENT ON COLUMN "admins" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "admins"
COMMENT ON COLUMN "admins" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "admins"
COMMENT ON COLUMN "admins" ."id" IS '';
-- reverse: create index "admins_phone_key" to table: "admins"
DROP INDEX "admins_phone_key";
-- reverse: create "admins" table
DROP TABLE "admins";
