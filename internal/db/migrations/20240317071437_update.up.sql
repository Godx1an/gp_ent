-- modify "school_fitness_test_items" table
ALTER TABLE "school_fitness_test_items" ADD COLUMN "school" character varying NOT NULL DEFAULT '', ADD COLUMN "item" character varying NOT NULL DEFAULT '';
