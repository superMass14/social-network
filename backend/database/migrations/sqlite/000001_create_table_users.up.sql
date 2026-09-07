-- database: ../../social_network.db
CREATE TABLE
	IF NOT EXISTS users (
		id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		user_name TEXT DEFAULT "",
		gender TEXT DEFAULT "",
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		user_type VARCHAR(25) NOT NULL,
		birth_date VARCHAR(12) NOT NULL,
		avatar VARCHAR(256),
		about_me TEXT DEFAULT ""
	);