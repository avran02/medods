CREATE TABLE "users" (
    "id" VARCHAR(255) PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE "refresh_tokens" (
    "user_id" VARCHAR(255) NOT NULL,
    "token_hash" VARCHAR(255) NOT NULL UNIQUE,
    "access_token_id" VARCHAR(255) NOT NULL UNIQUE,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    UNIQUE ("user_id")
);
