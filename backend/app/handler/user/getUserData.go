package user

import (
	"backend/models"
	"backend/app/cors"
	"backend/app/handler/groups"
	"backend/app/repositories"
	"backend/app/service"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// var userId int = 1

type RespUser struct {
	User       models.User    `json:"user"`
	Followers  []models.User  `json:"followers"`
	Groups     []models.Group `json:"groups"`
	OtherUsers []models.User  `json:"othersUsers"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var respUser RespUser
	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()
	session, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		log.Println("Invalid token", err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}

	userId := session.User_id

	respUser.User, _ = GetUserData(db, userId)
	idJoinedGroups, err1 := groups.GetJoinedGroups(db, userId)
	respUser.Groups, err = GetGroupData(db, idJoinedGroups)
	respUser.OtherUsers, _ = GetNonFollowedUsers(db, userId)
	// fmt.Println(respUser.OtherUsers)
	if err != nil || err1 != nil {
		fmt.Println("errr when getGroups")
		return
	}
	respUser.Followers, _ = GetFollowers(db, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respUser)
}

func GetGroupData(db *sql.DB, groupsId []int) ([]models.Group, error) {
	var groups []models.Group
	for _, idGroup := range groupsId {
		var group models.Group
		query := "SELECT * FROM groups WHERE id = ?"
		err := db.QueryRow(query, idGroup).Scan(&group.ID, &group.Name,
			&group.Description, &group.ID_User_Create, &group.Avatar,
			&group.Creation_Date)
		if err != nil {
			fmt.Println("errr ici")
			return []models.Group{}, fmt.Errorf("err scan user data: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

	func GetUserData(db *sql.DB, userID int) (models.User, error) {
		var user models.User
		query := "SELECT * FROM users WHERE id = ?"
		var user_name sql.NullString
		var gender sql.NullString
		var avatar sql.NullString
		var about_me sql.NullString
		err := db.QueryRow(query, userID).Scan(&user.Id, &user.First_name, &user.Last_name, &user_name, &gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &avatar, &about_me)
		if err != nil {
			log.Println("Error scanning row GetUserData: ", err.Error())
			return models.User{}, fmt.Errorf("err scan user data: %w", err)
		}
		user.User_name = repositories.GetStringValue(user_name)
		user.Gender = repositories.GetStringValue(gender)
		user.Avatar = repositories.GetStringValue(avatar)
		user.About_me = repositories.GetStringValue(about_me)
		return user, nil
	}

func GetFollowers(db *sql.DB, userID int) ([]models.User, error) {
	var followers []models.User
	query := "SELECT DISTINCT user_id_follower FROM users_followers WHERE user_id_followed = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		log.Println("Error querying row GetFollowers: ", err.Error())
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	var followerIDs []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}
		followerIDs = append(followerIDs, id)
	}

	query = "SELECT user_id_followed FROM users_followers WHERE user_id_follower = ?"
	rows, err = db.Query(query, userID)
	if err != nil {
		log.Println("Error querying row GetFollowers: ", err.Error())
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}

		if !contains(followerIDs, id) {
			followerIDs = append(followerIDs, id)
		}

	}

	for _, followerID := range followerIDs {
		user, err := GetUserData(db, followerID)
		if err != nil {
			return nil, fmt.Errorf("err execute all groups query: %w", err)
		}
		followers = append(followers, user)
	}

	return followers, nil
}
func GetNonFollowedUsers(db *sql.DB, userID int) ([]models.User, error) {
	var users []models.User
	query := `
        SELECT u.*
        FROM users u
        WHERE u.id NOT IN (
            SELECT user_id_follower
            FROM users_followers
            WHERE user_id_followed = ?
        ) AND u.id NOT IN (
            SELECT user_id_followed
            FROM users_followers
            WHERE user_id_follower = ?
        ) AND u.id != ?
    `
	rows, err := db.Query(query, userID, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &user.Avatar, &user.About_me)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func contains(arr []int, value int) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}
	return false
}
