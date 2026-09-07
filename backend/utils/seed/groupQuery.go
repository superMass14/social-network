package seed

import (
	"backend/models"
	"database/sql"
	"log"
)

// InsertMessage insère un nouveau message dans la base de données.
func InsertGroupMessage(db *sql.DB, userSender, groupReceiver int, messageContent, date string) error {

	query := `INSERT INTO messages (user_id_sender, group_id_receiver, message_content, date) VALUES (?, ?, ?, ?);`

	_, err := db.Exec(query, userSender, groupReceiver, messageContent, date)
	if err != nil {
		log.Println("Erreur lors de l'insertion du message group:", err)
		return err
	}

	return nil
}

// getGroupMessage récupère les messages pour un groupe spécifique
func GetGroupMessage(db *sql.DB, groupIDReceiver int) ([]models.Chat, error) {
	query := `SELECT * FROM messages WHERE group_id_receiver = ?`
	rows, err := db.Query(query, groupIDReceiver)
	if err != nil {
		log.Printf("Erreur lors de la récupération des messages du groupe: %v", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Chat

	for rows.Next() {
		var message models.Chat
		var userReceiver, groupIDReceiver *int
		err := rows.Scan(&message.ID, &message.UserSender, &userReceiver, &message.MessageContent, &groupIDReceiver, &message.Date)
		if err != nil {
			log.Printf("Erreur lors de la lecture des messages du groupe: %v", err)
			return nil, err
		}
		message.UserReceiver = userReceiver
		message.GroupIDReceiver = groupIDReceiver

		queryUser := `SELECT first_name FROM users WHERE id = ?`
		userRow, err := db.Query(queryUser, message.UserSender)
		if err != nil {
			log.Println("Erreur lors de la récupération du first_name:", err)
			return nil, err
		}
		defer userRow.Close()
		if userRow.Next() {
			err = userRow.Scan(&message.FirstName)
			if err != nil {
				log.Println("Erreur lors de la récupération du first_name:", err)
				return nil, err
			}
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Erreur lors de la récupération des messages du groupe: %v", err)
		return nil, err
	}

	return messages, nil
}

func GetLastGroupMessage(db *sql.DB, msg string) (models.Chat, error) {
	query := `SELECT * FROM messages WHERE message_content = ?`
	row, err := db.Query(query, msg)
	if err != nil {
		return models.Chat{}, err
	}
	var message models.Chat
	for row.Next() {
		var IDReceiver *int
		err = row.Scan(&message.ID, &message.UserSender, &IDReceiver, &message.MessageContent, &message.GroupIDReceiver, &message.Date)
		if err != nil {
			log.Println("Erreur lors de recuperer du message:", err)
			return models.Chat{}, err
		}
		// Récupérer le first_name de l'utilisateur qui a envoyé le message
		queryUser := `SELECT first_name FROM users WHERE id = ?`
		userRow, err := db.Query(queryUser, message.UserSender)
		if err != nil {
			log.Println("Erreur lors de la récupération du first_name:", err)
			return models.Chat{}, err
		}
		defer userRow.Close()
		if userRow.Next() {
			err = userRow.Scan(&message.FirstName)
			if err != nil {
				log.Println("Erreur lors de la récupération du first_name:", err)
				return models.Chat{}, err
			}
		}
	}
	return message, err
}

// GetGroup récupère les groupes liés à un utilisateur spécifique.
func GetGroup(db *sql.DB, userID int) ([]models.Group, error) {
	query := `
        SELECT groups.* FROM groups
        WHERE groups.id_user_create = ?
        OR groups.id IN (
            SELECT group_id FROM group_followers WHERE user_id = ?
        )
    `
	rows, err := db.Query(query, userID, userID)
	if err != nil {
		log.Printf("Erreur lors de la récupération des groupes: %v", err)
		return nil, err
	}
	defer rows.Close()
	var groups []models.Group
	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			log.Printf("Erreur lors de la lecture des groupes: %v", err)
			return nil, err
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Erreur lors de la récupération des groupes: %v", err)
		return nil, err
	}

	return groups, err
}

// UserEmail représente une structure contenant l'ID et l'email d'un utilisateur.
type UserEmail struct {
	UserID int
	Email  string
}

// GetFollowerGroup prend un ID de groupe et retourne un tableau d'IDs et emails d'utilisateurs qui suivent ce groupe.
func GetFollowerGroup(db *sql.DB, groupID int) ([]UserEmail, error) {
	query := `SELECT users.id, users.email FROM group_followers
	          INNER JOIN users ON group_followers.user_id = users.id
	          WHERE group_followers.group_id = ?`
	rows, err := db.Query(query, groupID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var userEmails []UserEmail
	for rows.Next() {
		var userEmail UserEmail
		if err := rows.Scan(&userEmail.UserID, &userEmail.Email); err != nil {
			log.Fatal(err)
			return nil, err
		}
		userEmails = append(userEmails, userEmail)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return userEmails, nil
}
