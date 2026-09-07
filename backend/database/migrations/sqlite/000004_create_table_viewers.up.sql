CREATE TABLE
    IF NOT EXISTS viewers (
        post_id TEXT NOT NULL,
        user_id INTEGER NOT NULL,
        FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE ON UPDATE CASCADE
    );