CREATE TABLE
    IF NOT EXISTS notifications (
        id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
        notification_type VARCHAR(25) NOT NULL,
        status TEXT NOT NULL,
        user_id_sender INTEGER NOT NULL,
        user_id_receiver INTEGER DEFAULT NULL,
        id_group INTEGER DEFAULT NULL,
        date DATETIME NOT NULL,
        FOREIGN KEY (id_group) REFERENCES groups (id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY (user_id_sender) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY (user_id_receiver) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
    );