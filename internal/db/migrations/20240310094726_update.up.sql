-- create "students" table
CREATE TABLE "students" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '请设置昵称', "nick_name" character varying NOT NULL DEFAULT '请设置昵称', "jpg_url" character varying NOT NULL DEFAULT '', PRIMARY KEY ("id"));
-- set comment to table: "students"
COMMENT ON TABLE "students" IS '用户表';
-- set comment to column: "id" on table: "students"
COMMENT ON COLUMN "students" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "students"
COMMENT ON COLUMN "students" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "students"
COMMENT ON COLUMN "students" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "students"
COMMENT ON COLUMN "students" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "students"
COMMENT ON COLUMN "students" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "students"
COMMENT ON COLUMN "students" ."deleted_at" IS '软删除时刻，带时区';
-- set comment to column: "name" on table: "students"
COMMENT ON COLUMN "students" ."name" IS '用户名';
-- set comment to column: "nick_name" on table: "students"
COMMENT ON COLUMN "students" ."nick_name" IS '用户昵称';
-- set comment to column: "jpg_url" on table: "students"
COMMENT ON COLUMN "students" ."jpg_url" IS '头像';
