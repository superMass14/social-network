CREATE TABLE
    IF NOT EXISTS posts (
        id TEXT UNIQUE PRIMARY KEY NOT NULL,
        content TEXT DEFAULT "NULL",
        media TEXT DEFAULT "",
        date TEXT NOT NULL,
        timestamp CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        user_id INTEGER NOT NULL,
        group_id INTEGER DEFAULT NULL,
        privacy TEXT DEFAULT "public",
        FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY ("group_id") REFERENCES "groups" ("id") ON DELETE CASCADE ON UPDATE CASCADE
    );