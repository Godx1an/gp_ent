-- create "schools" table
CREATE TABLE "schools" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "schools_name_key" to table: "schools"
CREATE UNIQUE INDEX "schools_name_key" ON "schools" ("name");
-- set comment to column: "id" on table: "schools"
COMMENT ON COLUMN "schools" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "schools"
COMMENT ON COLUMN "schools" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "schools"
COMMENT ON COLUMN "schools" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "schools"
COMMENT ON COLUMN "schools" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "schools"
COMMENT ON COLUMN "schools" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "schools"
COMMENT ON COLUMN "schools" ."deleted_at" IS '软删除时刻，带时区';
