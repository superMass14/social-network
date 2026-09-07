/* A ce niveau on aura besoin des deux table Users et UsersFollowers
   Pour etre dans la liste des online users d'un Users connecter par example
   ==> Soit le users te follow ou tu le follow (Cette liste sera afficher dans la partie des users avec qui il me chatter)
       Maintenant restera a determiner si le user est en ligne ou pas (Utilisation des sessions)

	   NB : Si lordre des users par rapport aux historique des messages va etre prise en compte faudra
	   joindre le table de messages aussi
*/

package seed

import (
	"backend/models"
	"database/sql"
	"log"
	"sort"
	"strings"
)

// Messages entre deux users

func SelectMsgBetweenUsers(db *sql.DB, CurrentUser, UserID int) ([]models.Chat, error) {
	query := `SELECT * FROM messages WHERE (user_id_sender = ? AND user_id_receiver = ? ) OR (user_id_sender = ? AND user_id_receiver = ?);`
	rows, err := db.Query(query, CurrentUser, UserID, UserID, CurrentUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chat []models.Chat

	for rows.Next() {
		var m models.Chat
		var groupIDReceiver *int
		// var userIDReceiver *int
		err := rows.Scan(&m.ID, &m.UserSender, &m.UserReceiver, &m.MessageContent, &m.GroupIDReceiver, &m.Date)
		if err != nil {
			return nil, err
		}
		m.GroupIDReceiver = groupIDReceiver
		row, _ := db.Query("SELECT first_name FROM users WHERE id = ?", m.UserSender)
		defer row.Close()

		for row.Next() {
			if err := row.Scan(&m.FirstName); err != nil {
				log.Println("error scanning first_name: ", err)
			}
		}
		chat = append(chat, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chat, nil
}

// SelectFollowersAndFollowing récupère tous les utilisateurs suivis et les suivants d'un utilisateur spécifique.
func SelectFollowersAndFollowing(db *sql.DB, userID int) ([]models.User, error) {
	// SQL query to retrieve followed and followers.
	query := `
		SELECT u.* FROM users u
		JOIN users_followers uf ON u.id = uf.user_id_followed OR u.id = uf.user_id_follower
		WHERE uf.user_id_follower = ? OR uf.user_id_followed = ?
	`

	rows, err := db.Query(query, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Use a map to keep track of unique users.
	userMap := make(map[int]models.User)
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &user.Avatar, &user.About_me)
		if err != nil {
			return nil, err
		}
		// Add the user to the map if it's not already present.
		if _, exists := userMap[user.Id]; !exists {
			userMap[user.Id] = user
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Convert the map to a slice.
	var users []models.User
	for _, user := range userMap {
		users = append(users, user)
	}

	return users, nil
}

// ListeUsers retourne un tableau de tableaux d'utilisateurs en se basant sur les deux fonctions
// SelectMsgBetweenUsers et SelectFollowersAndFollowing. Les utilisateurs sont triés par ordre des
// derniers messages s'il y en a, sinon par ordre alphabétique des noms des utilisateurs.
func ListeUsers(db *sql.DB, userID int) ([][]models.User, error) {
	// Récupérer tous les utilisateurs suivis et suivants.
	followersAndFollowing, err := SelectFollowersAndFollowing(db, userID)
	if err != nil {
		return nil, err
	}

	// Trier les utilisateurs par ordre alphabétique des noms de famille.
	sort.Slice(followersAndFollowing, func(i, j int) bool {
		return followersAndFollowing[i].First_name < followersAndFollowing[j].Last_name
	})

	// Récupérer les messages entre l'utilisateur actuel et chaque utilisateur suivi/suivi.
	var messages []models.Chat
	for _, user := range followersAndFollowing {
		userMessages, err := SelectMsgBetweenUsers(db, userID, user.Id)
		if err != nil {
			return nil, err
		}
		messages = append(messages, userMessages...)
	}

	// Si des messages ont été trouvés, trier les utilisateurs par ordre des derniers messages.
	if len(messages) > 0 {
		sort.Slice(followersAndFollowing, func(i, j int) bool {
			lastMessageI := findLastMessage(messages, followersAndFollowing[i].Id)
			lastMessageJ := findLastMessage(messages, followersAndFollowing[j].Id)

			// Trier par date du dernier message.
			return strings.Compare(lastMessageI.Date, lastMessageJ.Date) > 0
		})
	}

	// Grouper les utilisateurs par tableau de tableaux.
	var groupedUsers [][]models.User
	groupSize := 10 // Taille du groupe, ajustez selon vos besoins.
	for i := 0; i < len(followersAndFollowing); i += groupSize {
		end := i + groupSize
		if end > len(followersAndFollowing) {
			end = len(followersAndFollowing)
		}
		groupedUsers = append(groupedUsers, followersAndFollowing[i:end])
	}
	return groupedUsers, nil
}

// findLastMessage trouve le dernier message d'un utilisateur dans un tableau de messages.
func findLastMessage(messages []models.Chat, userID int) models.Chat {
	var lastMessage models.Chat
	for _, message := range messages {
		if message.UserSender == userID || *message.UserReceiver == userID {
			if strings.Compare(lastMessage.Date, message.Date) < 0 {
				lastMessage = message
			}
		}
	}
	return lastMessage
}

// InsertMessage insère un nouveau message dans la base de données.
func InsertMessage(db *sql.DB, userSender, userReceiver int, messageContent, date string) error {
	query := `INSERT INTO messages (user_id_sender, user_id_receiver, message_content, date) VALUES (?, ?, ?, ?);`
	_, err := db.Exec(query, userSender, userReceiver, messageContent, date)
	if err != nil {
		log.Println("Erreur lors de l'insertion du message:", err)
		return err
	}

	return nil
}

func GetLastMessage(db *sql.DB, msg string) (models.Chat, error) {
	query := `SELECT * FROM messages WHERE message_content = ?`
	row, err := db.Query(query, msg)
	if err != nil {
		return models.Chat{}, err
	}
	var message models.Chat
	for row.Next() {
		var groupIDReceiver *int
		err = row.Scan(&message.ID, &message.UserSender, &message.UserReceiver, &message.MessageContent, &groupIDReceiver, &message.Date)
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

func GetUserById(db *sql.DB, id int) (models.User, error) {
	query := `SELECT * FROM users WHERE id =?`
	row, err := db.Query(query, id)
	if err != nil {
		return models.User{}, nil
	}
	var user models.User
	for row.Next() {
		err = row.Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &user.Avatar, &user.About_me)
		if err != nil {
			log.Println("Erreur lors de la récupération des informations de l'utilisateur:", err)
			return models.User{}, nil
		}
	}
	return user, err
}
