CREATE TABLE
    IF NOT EXISTS sessions (
        token TEXT PRIMARY KEY NOT NULL,
        user_id INTEGER NOT NULL,
        expiration DATETIME NOT NULL,
        FOREIGN KEY ("user_id") REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
    )