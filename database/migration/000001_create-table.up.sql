CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY, 
    "username" VARCHAR(128) NOT NULL,
    "email" VARCHAR(128) UNIQUE NOT NULL,
    "password" VARCHAR(128) NOT NULL,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "photos" (
    "id" SERIAL PRIMARY KEY, 
    "title" VARCHAR(255) NOT NULL,
    "caption" VARCHAR(255) NOT NULL,
    "user_id" INT NOT NULL REFERENCES "users" (id) ON DELETE CASCADE,
    "photo_url" VARCHAR(255) NOT NULL
);
