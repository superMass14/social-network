CREATE TABLE
    IF NOT EXISTS event (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        group_id INTEGER NOT NULL,
        user_id INTEGER NOT NULL,
        description TEXT NOT NULL,
        event_date DATETIME NOT NULL,
        is_joined INTEGER DEFAULT 1,
        FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
    );