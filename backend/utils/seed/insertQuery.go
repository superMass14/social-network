package seed

import (
	"database/sql"
	"log"
)

func InsertData(db *sql.DB) {
	_, err := db.Exec(`
	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ('Madike', 'Yade', 'dickss', 'Male', 'dickss@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'private', '2000-01-01', 'http://localhost:8080/uploads/profilibg.jpg', 'about me...');

	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ('IBG', 'Gueye', 'ibg', 'Male', 'ibg@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'public', '2000-01-01', 'http://localhost:8080/uploads/profilibg.jpg', 'about me...');

	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ( 'Vincent', 'Ndour','vindour', 'Male', 'vindour@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'public', '2000-01-01', 'http://localhost:8080/uploads/profilibg.jpg', 'about me...');


	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ( 'Nikola', 'Faye', 'nixa','Male', 'nixa@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'private', '2000-01-01', 'http://localhost:8080/uploads/profilibg.jpg', 'about me...');

	INSERT INTO users (first_name, last_name, user_name, gender, email, password, user_type, birth_date, avatar, about_me)
	VALUES ( 'Daniella', 'Gueye', 'daniella', 'Female', 'dani@gmail.com', '$2a$10$rai5Jgvt7myDh.rltd.oseytjrOp3QRi9BDf7r0s133SW0HoOqewG', 'public', '2000-01-01', 'http://localhost:8080/uploads/profilibg.jpg', 'about me...');
	`)
	if err != nil {
		log.Fatal("Insert into users", err.Error())

	}

	//`user followers`

	_, err = db.Exec(`
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (1, 2);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (2, 1);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (1, 3);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (1, 4);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (2, 3);
		INSERT INTO users_followers (user_id_followed, user_id_follower)
		VALUES (2, 4);
	`)
	if err != nil {
		log.Fatal("Insert into users_followers", err.Error())
	}

	// insert into groups

	_, err = db.Exec(`
		INSERT INTO groups (name, description, id_user_create, avatar, creation_date)
		VALUES ('Group 1', 'first group', 1, 'http://localhost:8080/uploads/35a23cb2-d742-48fa-8d77-8a6e19bf571fsnk.jpg', '2024-03-05');
		INSERT INTO groups (name, description, id_user_create, avatar, creation_date)
		VALUES ('Group 2', 'second group', 3, 'http://localhost:8080/uploads/35a23cb2-d742-48fa-8d77-8a6e19bf571fsnk.jpg', '2024-03-05');
				`)
	if err != nil {
		log.Fatal("Insert into groups", err.Error())
	}

	// insert into message

	_, err = db.Exec(`
            INSERT INTO messages (user_id_sender, user_id_receiver,  message_content, date)
            VALUES (1, 2, 'This is the first message', '2024-03-05');
            INSERT INTO messages (user_id_sender, group_id_receiver,  message_content, date)
            VALUES (2, 1, 'This is the second message', '2024-03-06');
			INSERT INTO messages (user_id_sender, user_id_receiver,  message_content, date)
            VALUES (1, 2, 'This is the third message', '2024-03-06');
			`)

	if err != nil {
		log.Fatal("Insert into messages", err.Error())
	}

	// insert into group_followers

	_, err = db.Exec(`
            INSERT INTO group_followers (group_id, user_id)
            VALUES (1, 1);
            INSERT INTO group_followers (group_id, user_id)
            VALUES (1, 2);
			`)
	if err != nil {
		log.Fatal("Insert into group_followers", err.Error())
	}

	// insert into notifications

	_, err = db.Exec(
		`INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver, date)
		VALUES ("message", 0, 1, 3, '2024-04-13 04:23:09');
		INSERT INTO notifications (notification_type, status, user_id_sender, id_group, date)
		VALUES ("message", 0, 1, 2, '2024-04-13 04:23:09');
		INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver, date)
		VALUES ("follow", 0, 1, 4, '2024-04-13 04:23:09');
		`)
	if err != nil {
		log.Fatal("Insert into notifications", err.Error())
	}

	// insert into confirm

	_, err = db.Exec(
		`INSERT INTO confirm (user_id_asker, user_id_asked, group_id)
        VALUES (1, 2, NULL);

		INSERT INTO confirm (user_id_asker, user_id_asked, group_id)
        VALUES (2, NULL, 1);
        `)
	if err != nil {
		log.Fatal("Insert into confirm", err.Error())
	}

	

}
