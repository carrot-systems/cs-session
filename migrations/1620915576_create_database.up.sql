CREATE TABLE "public"."sessions" (
    "id" uuid unique NOT NULL,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "user_id" text,
    PRIMARY KEY ("id")
);
