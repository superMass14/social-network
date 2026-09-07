CREATE TABLE
    IF NOT EXISTS categories (
        post_id TEXT NOT NULL,
        category TEXT NOT NULL,
        FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE ON UPDATE CASCADE
    );