package joingroup

import (
	"backend/database"
	"backend/models"
	"backend/app/service"
	"database/sql"
	"fmt"
	"net/http"
)

func NotificationUserData(r *http.Request) ([]models.NotificationUser, error) {
	id_user, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		fmt.Println("Error verifying ", err.Error())
		return nil, err
	}
	db := database.NewDatabase()
	stm := `
	SELECT notifications.id, notifications.notification_type, 
	notifications.status, notifications.user_id_sender, 
	notifications.user_id_receiver, notifications.id_group, 
	notifications.date, users.first_name, users.last_name, 
	users.user_name, users.email
	FROM notifications
	JOIN users ON notifications.user_id_sender = users.id
	WHERE notifications.user_id_receiver = ? AND notifications.status = ?
	`
	query, err := db.Prepare(stm)
	if err != nil {
		fmt.Println("Erreur lors de la preparation de la requete ", err)
		return []models.NotificationUser{}, err
	}
	defer query.Close()

	rows, err := query.Query(id_user.User_id, 0)
	if err != nil {
		fmt.Println("Erreur Query ", err.Error())
	}
	defer rows.Close()
	var notifications []models.NotificationUser
	for rows.Next() {
		var notification models.NotificationUser
		var groupId sql.NullInt64
		err := rows.Scan(&notification.ID, &notification.Type, &notification.Status, &notification.UserIDSender, &notification.UserIDReceveived, &groupId, &notification.Date, &notification.First_name, &notification.Last_name, &notification.User_name, &notification.Email)
		notification.GroupId = getIntValue(groupId)
		if err != nil {
			fmt.Println("Error scanning row: ", err.Error())
			return []models.NotificationUser{}, err
		}
		notifications = append(notifications, notification)

	}
	return notifications, nil
}
