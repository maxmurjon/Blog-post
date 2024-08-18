CREATE TABLE IF NOT EXISTS "comments"(
    "id" SERIAL PRIMARY KEY,
    "post_id" INTEGER REFERENCES "posts"("id"),
    "users_id" INTEGER REFERENCES "users"("id"),
    "content" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)