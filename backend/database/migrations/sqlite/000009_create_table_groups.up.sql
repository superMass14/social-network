CREATE TABLE
    IF NOT EXISTS groups (
        id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
        name varchar(256) NOT NULL,
        description text NOT NULL,
        id_user_create integer NOT NULL,
        avatar VARCHAR(256),
        creation_date VARCHAR(256) NOT NULL,
        FOREIGN KEY (id_user_create) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
    );