-- create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "name" character varying NOT NULL DEFAULT '请设置昵称', "nick_name" character varying NOT NULL DEFAULT '请设置昵称', "jpg_url" character varying NOT NULL DEFAULT '', "phone" character varying NOT NULL DEFAULT '', "password" character varying NOT NULL DEFAULT '', "is_frozen" boolean NOT NULL DEFAULT false, "is_recharge" boolean NOT NULL DEFAULT false, "user_type" character varying NOT NULL DEFAULT 'personal', "pop_version" character varying NOT NULL DEFAULT '', "area_code" character varying NOT NULL DEFAULT '+86', "email" character varying NOT NULL DEFAULT '', "cloud_space" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to table: "users"
COMMENT ON TABLE "users" IS '用户表';
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
-- set comment to column: "name" on table: "users"
COMMENT ON COLUMN "users" ."name" IS '用户名';
-- set comment to column: "nick_name" on table: "users"
COMMENT ON COLUMN "users" ."nick_name" IS '用户昵称';
-- set comment to column: "jpg_url" on table: "users"
COMMENT ON COLUMN "users" ."jpg_url" IS '头像';
-- set comment to column: "phone" on table: "users"
COMMENT ON COLUMN "users" ."phone" IS '用户的手机号';
-- set comment to column: "password" on table: "users"
COMMENT ON COLUMN "users" ."password" IS '密码';
-- set comment to column: "is_frozen" on table: "users"
COMMENT ON COLUMN "users" ."is_frozen" IS '是否冻结';
-- set comment to column: "is_recharge" on table: "users"
COMMENT ON COLUMN "users" ."is_recharge" IS '是否充值过';
-- set comment to column: "user_type" on table: "users"
COMMENT ON COLUMN "users" ."user_type" IS '用户类型';
-- set comment to column: "pop_version" on table: "users"
COMMENT ON COLUMN "users" ."pop_version" IS '用户最新弹窗版本';
-- set comment to column: "area_code" on table: "users"
COMMENT ON COLUMN "users" ."area_code" IS '国家区号';
-- set comment to column: "email" on table: "users"
COMMENT ON COLUMN "users" ."email" IS '邮箱';
-- set comment to column: "cloud_space" on table: "users"
COMMENT ON COLUMN "users" ."cloud_space" IS '云盘空间';
