-- create "admins" table
CREATE TABLE "admins" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "phone" character varying NOT NULL, "nickname" character varying NOT NULL, "password" character varying NOT NULL, "school" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- create index "admins_phone_key" to table: "admins"
CREATE UNIQUE INDEX "admins_phone_key" ON "admins" ("phone");
-- set comment to column: "id" on table: "admins"
COMMENT ON COLUMN "admins" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "admins"
COMMENT ON COLUMN "admins" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "admins"
COMMENT ON COLUMN "admins" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "admins"
COMMENT ON COLUMN "admins" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "admins"
COMMENT ON COLUMN "admins" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "admins"
COMMENT ON COLUMN "admins" ."deleted_at" IS '软删除时刻，带时区';
