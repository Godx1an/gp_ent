-- create "fitness_test_items" table
CREATE TABLE "fitness_test_items" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "item" character varying NOT NULL, PRIMARY KEY ("id"));
-- create index "fitness_test_items_item_key" to table: "fitness_test_items"
CREATE UNIQUE INDEX "fitness_test_items_item_key" ON "fitness_test_items" ("item");
-- set comment to column: "id" on table: "fitness_test_items"
COMMENT ON COLUMN "fitness_test_items" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "fitness_test_items"
COMMENT ON COLUMN "fitness_test_items" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "fitness_test_items"
COMMENT ON COLUMN "fitness_test_items" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "fitness_test_items"
COMMENT ON COLUMN "fitness_test_items" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "fitness_test_items"
COMMENT ON COLUMN "fitness_test_items" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "fitness_test_items"
COMMENT ON COLUMN "fitness_test_items" ."deleted_at" IS '软删除时刻，带时区';
-- create "school_fitness_test_items" table
CREATE TABLE "school_fitness_test_items" ("id" bigint NOT NULL, "created_by" bigint NOT NULL DEFAULT 0, "updated_by" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NOT NULL, "max_participants" bigint NOT NULL DEFAULT 0, "avg_time_per_person" bigint NOT NULL DEFAULT 0, "school_id" bigint NOT NULL DEFAULT 0, "item_id" bigint NOT NULL DEFAULT 0, PRIMARY KEY ("id"));
-- set comment to column: "id" on table: "school_fitness_test_items"
COMMENT ON COLUMN "school_fitness_test_items" ."id" IS '19 位雪花 ID';
-- set comment to column: "created_by" on table: "school_fitness_test_items"
COMMENT ON COLUMN "school_fitness_test_items" ."created_by" IS '创建者 ID';
-- set comment to column: "updated_by" on table: "school_fitness_test_items"
COMMENT ON COLUMN "school_fitness_test_items" ."updated_by" IS '更新者 ID';
-- set comment to column: "created_at" on table: "school_fitness_test_items"
COMMENT ON COLUMN "school_fitness_test_items" ."created_at" IS '创建时刻，带时区';
-- set comment to column: "updated_at" on table: "school_fitness_test_items"
COMMENT ON COLUMN "school_fitness_test_items" ."updated_at" IS '更新时刻，带时区';
-- set comment to column: "deleted_at" on table: "school_fitness_test_items"
COMMENT ON COLUMN "school_fitness_test_items" ."deleted_at" IS '软删除时刻，带时区';
