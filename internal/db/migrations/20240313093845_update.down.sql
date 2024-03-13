-- reverse: set comment to column: "deleted_at" on table: "schools"
COMMENT ON COLUMN "schools" ."deleted_at" IS '';
-- reverse: set comment to column: "updated_at" on table: "schools"
COMMENT ON COLUMN "schools" ."updated_at" IS '';
-- reverse: set comment to column: "created_at" on table: "schools"
COMMENT ON COLUMN "schools" ."created_at" IS '';
-- reverse: set comment to column: "updated_by" on table: "schools"
COMMENT ON COLUMN "schools" ."updated_by" IS '';
-- reverse: set comment to column: "created_by" on table: "schools"
COMMENT ON COLUMN "schools" ."created_by" IS '';
-- reverse: set comment to column: "id" on table: "schools"
COMMENT ON COLUMN "schools" ."id" IS '';
-- reverse: create index "schools_name_key" to table: "schools"
DROP INDEX "schools_name_key";
-- reverse: create "schools" table
DROP TABLE "schools";
