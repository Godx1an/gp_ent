-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "phone" character varying NOT NULL, "nickname" character varying NOT NULL, "password" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "users_phone_key" to table: "users"
CREATE UNIQUE INDEX "users_phone_key" ON "users" ("phone");
-- set comment to column: "id" on table: "users"
COMMENT ON COLUMN "users" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "users"
COMMENT ON COLUMN "users" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "users"
COMMENT ON COLUMN "users" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "users"
COMMENT ON COLUMN "users" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "users"
COMMENT ON COLUMN "users" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "users"
COMMENT ON COLUMN "users" ."deleted_at" IS '软删除时刻，带时区';
