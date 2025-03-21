-- -------------------------------------------------------------
-- TablePlus 6.2.1(578)
--
-- https://tableplus.com/
--
-- Database: kreditplus
-- Generation Time: 2025-03-21 21:47:47.3260
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."limits";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS limits_id_seq;

-- Table Definition
CREATE TABLE "public"."limits" (
    "id" int8 NOT NULL DEFAULT nextval('limits_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "tenor_months" int8,
    "limit_amount" numeric,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."payment_schedules";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS payment_schedules_id_seq;

-- Table Definition
CREATE TABLE "public"."payment_schedules" (
    "id" int8 NOT NULL DEFAULT nextval('payment_schedules_id_seq'::regclass),
    "transaction_id" int8 NOT NULL,
    "due_date" timestamptz,
    "amount" numeric,
    "status" text,
    "payment_date" timestamptz,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."transactions";
-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS transactions_id_seq;

-- Table Definition
CREATE TABLE "public"."transactions" (
    "id" int8 NOT NULL DEFAULT nextval('transactions_id_seq'::regclass),
    "user_id" int8 NOT NULL,
    "contract_number" text NOT NULL,
    "otr" numeric,
    "admin_fee" numeric,
    "installment_amount" numeric,
    "interest_amount" numeric,
    "asset_name" text,
    "tenor_months" int8,
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
    "nik" text NOT NULL,
    "full_name" text,
    "legal_name" text,
    "birth_place" text,
    "birth_date" date,
    "salary" numeric,
    "photo_id_card" text,
    "photo_selfie" text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."limits" ("id", "user_id", "tenor_months", "limit_amount", "created_at", "updated_at", "deleted_at") VALUES
(1, 1, 3, 5000000, '2025-03-21 21:38:01.819987+07', '2025-03-21 21:39:41.255403+07', NULL);

INSERT INTO "public"."payment_schedules" ("id", "transaction_id", "due_date", "amount", "status", "payment_date", "created_at", "updated_at", "deleted_at") VALUES
(1, 1, '2025-04-21 21:38:04.707068+07', 1100000, 'PAID', '2025-03-25 07:00:00+07', '2025-03-21 21:38:04.707115+07', '2025-03-21 21:38:18.833443+07', NULL),
(2, 1, '2025-05-21 21:38:04.70949+07', 1100000, 'PAID', '2025-04-25 07:00:00+07', '2025-03-21 21:38:04.709524+07', '2025-03-21 21:39:36.090114+07', NULL),
(3, 1, '2025-06-21 21:38:04.710086+07', 1100000, 'PAID', '2025-05-25 07:00:00+07', '2025-03-21 21:38:04.710107+07', '2025-03-21 21:39:41.245549+07', NULL);

INSERT INTO "public"."transactions" ("id", "user_id", "contract_number", "otr", "admin_fee", "installment_amount", "interest_amount", "asset_name", "tenor_months", "created_at", "updated_at", "deleted_at") VALUES
(1, 1, 'TX-001', 3000000, 100000, 1100000, 200000, 'Yamaha NMax', 3, '2025-03-21 21:38:04.703133+07', '2025-03-21 21:38:04.703133+07', NULL);

INSERT INTO "public"."users" ("id", "nik", "full_name", "legal_name", "birth_place", "birth_date", "salary", "photo_id_card", "photo_selfie", "created_at", "updated_at", "deleted_at") VALUES
(1, '1234567891', 'John Doe', 'Johnathan Doe', 'Jakarta', '1995-05-01', 5000000, '', '', '2025-03-21 21:37:58.586469+07', '2025-03-21 21:37:58.586469+07', NULL);

ALTER TABLE "public"."limits" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE SET NULL ON UPDATE CASCADE;


-- Indices
CREATE INDEX idx_limits_deleted_at ON public.limits USING btree (deleted_at);


-- Indices
CREATE INDEX idx_payment_schedules_deleted_at ON public.payment_schedules USING btree (deleted_at);


-- Indices
CREATE UNIQUE INDEX uni_transactions_contract_number ON public.transactions USING btree (contract_number);
CREATE INDEX idx_transactions_deleted_at ON public.transactions USING btree (deleted_at);


-- Indices
CREATE UNIQUE INDEX uni_users_nik ON public.users USING btree (nik);
CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);
