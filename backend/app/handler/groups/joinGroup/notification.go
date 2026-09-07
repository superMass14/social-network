package joingroup

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func RecupeIdAdminGroup(idGroup int, db *sql.DB) (int, error) {
	req := `
		SELECT id_user_create FROM groups WHERE id = ?
	`

	stm, err := db.Prepare(req)
	if err != nil {
		return 0, err
	}

	defer stm.Close()

	var id_user_greate_group int

	err = stm.QueryRow(idGroup).Scan(&id_user_greate_group)
	if err != nil {
		return 0, err
	}
	return id_user_greate_group, nil
}

func InsertNotification(idGroup int, notification_type string, user_id_sender, id_user_receiver int, db *sql.DB) error {
	id_user_created_group, err := RecupeIdAdminGroup(idGroup, db)
	if notification_type == "SuggestFriend" {
		id_user_created_group = id_user_receiver
		// user_id_sender, err = RecupeIdAdminGroup(idGroup, db)
	}
	if err != nil {
		log.Fatal("Erreur lors de la recuperation de l'id de l'admin group ", err)
	}
	checkNotif, _ := CheckNotifAndType(db, idGroup, user_id_sender, notification_type)
	if !checkNotif {

		req := `
			INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver, id_group, date) VALUES ($1, $2, $3, $4, $5, $6)
		`
		stm, err := db.Prepare(req)

		if err != nil {
			return err
		}

		defer stm.Close()

		date := time.Now()
		formattedDateTime := date.Format("2006-01-02 15:04:05")
		_, err = stm.Exec(notification_type, 0, user_id_sender, id_user_created_group, idGroup, formattedDateTime)
		if err != nil {
			fmt.Println("Erreur lors de l'execution de la requete inserte dans la base ", err)
			return err
		}
	} else {
		fmt.Println("Vous avez une demande en cours")
	}
	return nil
}

func CheckJoinNotification(id_user_created_group int, id_user_connected int, idGroup int, db *sql.DB) (int, int) {
	req := `
        SELECT id, status FROM notifications WHERE user_id_sender = ? AND user_id_receiver = ? AND id_group = ?
    `
	stm, err := db.Prepare(req)
	if err != nil {
		fmt.Println("Error preparing request checkJoinNotification: ", err)
		return 0, 0
	}

	defer stm.Close()

	var id_notification int
	var state int

	err = stm.QueryRow(id_user_connected, id_user_created_group, idGroup).Scan(&id_notification, &state)
	if err != nil {
		// fmt.Println("Error querying checkJoinNotification: ", err)
		return 0, 0
	}
	return id_notification, state
}

func CheckNotifAndType(db *sql.DB, groupID int, userID int, notificationType string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM notifications WHERE notification_type = $1 AND user_id_sender = $2 AND id_group = $3", notificationType, userID, groupID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AcceptOrNo(db *sql.DB, user_id_sender, user_id_receiver, id_group int, val string) bool {
	stm := `
		UPDATE notifications SET status = ? WHERE user_id_sender = ? AND user_id_receiver = ? AND id_group = ?
	`
	query, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête: ", err)
		return false
	}
	defer query.Close() // Assurez-vous de fermer la requête préparée après son utilisation

	// Exécute la requête préparée avec les paramètres spécifiques
	_, err = query.Exec(val, user_id_sender, user_id_receiver, id_group) // '1' est la nouvelle valeur de 'status'
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête: ", err)
		return false
	}

	return true
}

func AcceptOrNoSugg(db *sql.DB, user_id_sender, user_id_receiver int, val string) bool {
	stm := `
		UPDATE notifications SET status = ? WHERE user_id_sender = ? AND user_id_receiver = ? AND notification_type = ?
	`
	query, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête: ", err)
		return false
	}
	defer query.Close() // Assurez-vous de fermer la requête préparée après son utilisation

	// Exécute la requête préparée avec les paramètres spécifiques
	_, err = query.Exec(val, user_id_sender, user_id_receiver, "follow") // '1' est la nouvelle valeur de 'status'
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête: ", err)
		return false
	}

	return true
}

func InsertGroupFollowers(db *sql.DB, user_id, group_id int) {
	isExist, errExist := VerifyGroupFollow(db, user_id, group_id)
	if errExist != nil {
		log.Println("error in insertGroupFollower: ", errExist)
		return
	}

	if !isExist {

		stm := `
		INSERT INTO group_followers VALUES (?, ?)
	`
		query, err := db.Prepare(stm)
		if err != nil {
			fmt.Println("Erreur lors du prepation de la requete ", err)
			return
		}
		defer query.Close()

		_, err = query.Exec(user_id, group_id)
		if err != nil {
			fmt.Println("Erreur lors de l'execution de la requete ", err)
		}
	}
}
func VerifyGroupFollow(DB *sql.DB, user_id, group_id int) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM group_followers WHERE user_id = $1 AND group_id = $2", user_id, group_id).Scan(&count)
	if err != nil {
		fmt.Println("Erreur requete ", err)
		return false, err
	}
	return count > 0, nil
}

func InsertFollowNotification(followerId, followedId int, db *sql.DB) error {
	// Insertion de la notification
	fmt.Println("Insertion de la notification")
	checkNotif, err := CheckNotifExist(db, followerId, followedId)
	if err != nil {
		fmt.Println("Erreur checked ", err.Error())
		return err
	}
	fmt.Println("CheckNotifExist ", checkNotif)
	if !checkNotif {
		check, state := CheckJoinFollowNotification(followerId, followedId, db)
		if check == 0 && state == 0 || check > 0 && state != -1 {
			req := `
			INSERT INTO notifications (notification_type, status, user_id_sender, user_id_receiver, date) VALUES ($1, $2, $3, $4, $5)
		`
			stm, err := db.Prepare(req)

			if err != nil {
				fmt.Println(err)
				return err
			}

			defer stm.Close()
			date := time.Now()
			_, err = stm.Exec("follow", 0, followerId, followedId, date)
			if err != nil {
				fmt.Println("Erreur lors de l'execution de la requete inserte dans la base ", err)
				return err
			}
		} else {
			fmt.Println("Vous avez une demande en cours")
			return err
		}
	}
	return nil
}

func CheckJoinFollowNotification(followerId, followedId int, db *sql.DB) (int, int) {
	fmt.Printf("FollwerId = %v, FoolowedId = %v", followerId, followedId)

	req := `
        SELECT id, status FROM notifications WHERE user_id_sender = $1 AND user_id_receiver = $2
    `

	stm, err := db.Prepare(req)
	if err != nil {
		fmt.Println("Error preparing request checkJoinNotification: ", err)
		return 0, 0
	}

	defer stm.Close()

	var id_notification int
	var state int

	err = stm.QueryRow(followerId, followedId).Scan(&id_notification, &state)
	if err != nil {
		fmt.Println("Error querying checkJoinNotification: ", err)
		return 0, 0
	}

	fmt.Println("Il existe dejà une demande de follow")
	return id_notification, state
}

func CheckNotifExist(db *sql.DB, followerId, followedId int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM notifications WHERE user_id_sender = $1 AND user_id_receiver = $2 AND notification_type = $3", followerId, followedId, "follow").Scan(&count)
	if err != nil {
		fmt.Println("Erreur requete ", err)
		return false, err
	}
	fmt.Println("Count = ", count)
	return count > 0, nil
}
func DeleteNotification(db *sql.DB, val string) {
	stm := `
		DELETE FROM notifications WHERE status = $1
	`
	query, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Erreur lors de la préparation de la requête: ", err)
		return
	}
	defer query.Close() // Assurez-vous de fermer la requête préparée après son utilisation
	// Exécute la requête préparée avec les paramètres spécifiques
	_, err = query.Exec(val) // '1' est la nouvelle valeur de 'status'
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête: ", err)
		return
	}
}
