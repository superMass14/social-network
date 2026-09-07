CREATE TABLE
    IF NOT EXISTS messages (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        user_id_sender INTEGER NOT NULL,
        user_id_receiver INTERGER DEFAULT NULL,
        message_content TEXT NOT NULL,
        group_id_receiver INTERGER DEFAULT NULL,
        date TEXT,
        FOREIGN KEY (user_id_sender) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY (user_id_receiver) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
    );