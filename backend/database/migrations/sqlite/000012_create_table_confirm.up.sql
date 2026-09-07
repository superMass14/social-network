CREATE TABLE
    IF NOT EXISTS confirm (
        user_id_asker INTEGER NOT NULL,
        user_id_asked INTEGER DEFAULT "NULL",
        group_id INTEGER DEFAULT "NULL",
        FOREIGN KEY (user_id_asker) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY (user_id_asked) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE
    );