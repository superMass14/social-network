CREATE TABLE
    IF NOT EXISTS users_followers (
        user_id_followed INTEGER NOT NULL,
        user_id_follower INTEGER NOT NULL,
        FOREIGN KEY (user_id_followed) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY (user_id_follower) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
    );