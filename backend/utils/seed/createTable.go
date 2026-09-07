package seed

import (
	"database/sql"
	"log"
)

func CreateTable(db *sql.DB) {

	// creation de la table user
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		    id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT NOT NULL,
			first_name TEXT NOT NULL,
			last_name TEXT NOT NULL,
			user_name TEXT,
			gender TEXT DEFAULT "",
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			user_type VARCHAR(25) NOT NULL,
			birth_date VARCHAR(12) NOT NULL,
			avatar VARCHAR(256),
			about_me TEXT
			)
			 `)
	if err != nil {
		log.Fatal("User table ", err.Error())
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users_followers (
		user_id_followed INTEGER NOT NULL,
		user_id_follower INTEGER NOT NULL,
        FOREIGN KEY(user_id_followed) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
        FOREIGN KEY(user_id_follower) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
			)
			 `)
	if err != nil {
		log.Fatal("Users_followers table", err.Error())
	}

	//? Creation de la table posts
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id TEXT UNIQUE PRIMARY KEY NOT NULL,
			content TEXT DEFAULT "NULL",
			media TEXT DEFAULT "",
			date TEXT NOT NULL,
			timestamp CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_id INTEGER NOT NULL,
			group_id INTEGER DEFAULT NULL,
			privacy TEXT DEFAULT "public",
			FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
			FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE ON UPDATE CASCADE
		 )
		`)
	if err != nil {
		log.Fatal("Posts table ", err.Error())
	}

	//? Creation de la table viewers
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS viewers (
			post_id TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY("user_id") REFERENCES "users"("id") ON DELETE CASCADE ON UPDATE CASCADE
			FOREIGN KEY("post_id") REFERENCES "posts"("id") ON DELETE CASCADE ON UPDATE CASCADE
		 )
		`)
	if err != nil {
		log.Fatal("Viewers table ", err.Error())
	}

	// Création tavle belong
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			post_id TEXT NOT NULL,
			category TEXT NOT NULL,
			FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE
		)
		`)
	if err != nil {
		log.Fatal("PostCategory table ", err.Error())
	}

	//? Créate de la table comment
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comment (
			id INTEGER UNIQUE PRIMARY KEY AUTOINCREMENT NOT NULL,
			content TEXT DEFAULT "NULL",
			date TEXT NOT NULL,
			media TEXT DEFAULT NULL,
			post_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY(post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)
		`)
	if err != nil {
		log.Fatal("Comment table", err.Error())
	}

	// Creation de la table session
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS sessions (
    	token TEXT PRIMARY KEY NOT NULL,
    	user_id INTEGER NOT NULL,
    	expiration DATETIME NOT NULL,
		FOREIGN KEY("user_id") REFERENCES  users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)
		`)
	if err != nil {
		log.Fatal("Sessions table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS messages (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			user_id_sender INTEGER NOT NULL,
			user_id_receiver INTERGER DEFAULT NULL,
			message_content TEXT NOT NULL,
			group_id_receiver INTERGER DEFAULT NULL,
			date TEXT,
			FOREIGN KEY (user_id_sender) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY (user_id_receiver) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("Message table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS groups (
			id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
			name varchar(256)NOT NULL,
			description text NOT NULL,
			id_user_create integer NOT NULL,
			avatar VARCHAR(256),
			creation_date VARCHAR(256) NOT NULL,
			FOREIGN KEY(id_user_create) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("Groups table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS group_followers (
			user_id INTEGER NOT NULL,
			group_id INTEGER NOT NULL,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(group_id) REFERENCES groups(id) ON DELETE CASCADE ON UPDATE CASCADE
			)`)
	if err != nil {
		log.Fatal("group_followers table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS notifications (
			id INTEGER NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
			notification_type VARCHAR(25) NOT NULL,
			status NOT NULL,
			user_id_sender INTEGER NOT NULL,
			user_id_receiver INTEGER DEFAULT NULL,
			id_group INTEGER DEFAULT NULL,
			date DATETIME NOT NULL,
            FOREIGN KEY(id_group) REFERENCES groups(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(user_id_sender) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(user_id_receiver) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("Notifications table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS confirm (
			user_id_asker INTEGER NOT NULL,
			user_id_asked INTEGER DEFAULT "NULL",
			group_id INTEGER DEFAULT "NULL",
            FOREIGN KEY(user_id_asker) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(user_id_asked) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
            FOREIGN KEY(group_id) REFERENCES groups(id) ON DELETE CASCADE ON UPDATE CASCADE

		)`)
	if err != nil {
		log.Fatal("Confirm table", err.Error())
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "event" (
			"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			"group_id" INTEGER NOT NULL,
			"user_id" INTEGER NOT NULL,
			description TEXT NOT NULL,
			event_date DATETIME NOT NULL, "is_joined" INTEGER DEFAULT 1,
			FOREIGN KEY("group_id") REFERENCES groups(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY("user_id") REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("event table", err.Error())
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS  event_joined (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			group_id INTEGER REFERENCES groups("id") ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(event_id) REFERENCES event("id") ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("event_joined table", err.Error())
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS event_notjoined (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			group_id INTEGER NOT NULL,
			FOREIGN KEY(event_id) REFERENCES event(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(group_id) REFERENCES groups(id) ON DELETE CASCADE ON UPDATE CASCADE
		)`)
	if err != nil {
		log.Fatal("event_joined table", err.Error())
	}

}
