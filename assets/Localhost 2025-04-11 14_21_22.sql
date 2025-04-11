-- -------------------------------------------------------------
-- TablePlus 6.4.0(598)
--
-- https://tableplus.com/
--
-- Database: dgw_test
-- Generation Time: 2025-04-11 14:21:33.9710
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."categories";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS categories_id_seq;

-- Table Definition
CREATE TABLE "public"."categories" (
    "id" int8 NOT NULL DEFAULT nextval('categories_id_seq'::regclass),
    "name" text NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."products";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS products_id_seq;

-- Table Definition
CREATE TABLE "public"."products" (
    "id" int8 NOT NULL DEFAULT nextval('products_id_seq'::regclass),
    "name" text NOT NULL,
    "price" numeric NOT NULL,
    "category_id" int8 NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "name" text NOT NULL,
    "username" text NOT NULL,
    "password" text NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."categories" ("id", "name", "created_at", "updated_at", "deleted_at") VALUES
(1, 'Phone', '2025-04-11 13:03:05.36292+07', '2025-04-11 13:03:05.36292+07', NULL),
(2, 'Laptop', '2025-04-11 13:03:26.157447+07', '2025-04-11 13:03:26.157447+07', NULL),
(3, 'Tablet 2', '2025-04-11 13:03:32.742285+07', '2025-04-11 13:04:26.417351+07', '2025-04-11 13:04:49.949997+07');

INSERT INTO "public"."products" ("id", "name", "price", "category_id", "created_at", "updated_at", "deleted_at") VALUES
(1, 'Iphone', 1000, 1, '2025-04-11 13:09:08.241206+07', '2025-04-11 13:09:08.241206+07', NULL),
(2, 'Samsung', 1100, 1, '2025-04-11 13:09:22.531176+07', '2025-04-11 13:09:22.531176+07', NULL),
(3, 'Nokia', 100, 1, '2025-04-11 13:09:28.897941+07', '2025-04-11 13:09:28.897941+07', NULL),
(4, 'Macbook', 1000, 2, '2025-04-11 13:09:39.44494+07', '2025-04-11 13:09:39.44494+07', NULL),
(5, 'ASUS', 1200, 2, '2025-04-11 13:09:47.188954+07', '2025-04-11 13:09:47.188954+07', NULL),
(6, 'Legion', 1200, 2, '2025-04-11 13:09:56.954792+07', '2025-04-11 13:09:56.954792+07', NULL),
(7, 'Lenovo Phone', 400, 1, '2025-04-11 13:16:15.741724+07', '2025-04-11 13:29:12.822933+07', '2025-04-11 13:29:46.399719+07');

INSERT INTO "public"."users" ("id", "name", "username", "password", "created_at", "updated_at", "deleted_at") VALUES
(1, 'Admin 1', 'admin', '$2a$14$sxMJ9SWtzww//oaHmEVJRuOWGpF4gDRLtByrbhBvesw.g6cPMDWrm', '2025-04-11 10:24:12.029226+07', '2025-04-11 10:24:12.029226+07', NULL),
(2, 'Admin 2', 'admin2', '$2a$14$6isqE8XcZ3hNPn7ysLCG/.a/BOlU89Ys3q6yUVoYn5pUcDlQOv3Da', '2025-04-11 10:25:24.749197+07', '2025-04-11 10:25:24.749197+07', NULL),
(3, 'John Doe2', 'johndoe2', '$2a$14$6isqE8XcZ3hNPn7ysLCG/.a/BOlU89Ys3q6yUVoYn5pUcDlQOv3Da', '2025-04-11 10:29:21.108432+07', '2025-04-11 10:57:24.389675+07', '2025-04-11 10:58:28.044598+07');



-- Indices
CREATE UNIQUE INDEX uni_categories_name ON public.categories USING btree (name);
CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);
ALTER TABLE "public"."products" ADD FOREIGN KEY ("category_id") REFERENCES "public"."categories"("id");


-- Indices
CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


-- Indices
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);
CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);
