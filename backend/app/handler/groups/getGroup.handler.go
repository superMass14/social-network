package groups

import (
	"backend/app/cors"
	"backend/app/service"
	"backend/models"
	utils "backend/utils"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	ID       int              `json:"id"`
	Content  string           `json:"content"`
	Media    string           `json:"media"`
	Date     string           `json:"date"`
	User_id  int              `json:"userId"`
	FullName string           `json:"fullname"`
	Username string           `json:"username"`
	User     models.User      `json:"user"`
	Group_id int              `json:"groupId"`
	Privacy  string           `json:"privacy"`
	Comments []models.Comment `json:"comments"`
	Likes    int              `json:"likes"`
	Dislikes int              `json:"dislikes"`
}

type ResponseGroup struct {
	GroupData   models.Group      `json:"group_data"`
	Posts       []models.DataPost `json:"posts"`
	Members     []models.User     `json:"members"`
	Followers   []models.User     `json:"followers"`
	Event       []models.Event    `json:"events"`
	EventJoined []models.Event    `json:"events_joined"`
}

// const userId int = 1

func GetGroup(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()

	groupId, err := GetGroupIDFromRequest(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	session, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		log.Println("Invalid Token ", err)
		utils.Alert(w, models.Errormessage{
			Type:       "Get group",
			Msg:        "Invalid Token",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	userId := session.User_id

	isExist, err := CheckIfFollowerExists(db, groupId, userId)

	if !isExist {

		log.Println("Not allowed ")
		utils.Alert(w, models.Errormessage{
			Type:       "Get group",
			Msg:        "Not allowed",
			StatusCode: http.StatusMethodNotAllowed,
		})
		return
	}

	allPosts, err := service.PostServ.FetchPostGroup(groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	members, err := getMembers(db, groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	groupData, err := getInfoGroup(db, groupId)
	if err != nil {
		fmt.Println(err)
		return
	}

	followers, err := GetFollowers(db, userId)

	membersFilt := filterNonMembers(followers, members)
	if err != nil {
		fmt.Println(err)
		return
	}
	// events, eventJoined, err := GetEvent(db, groupId, userId)
	eventJoined, err := GetJoinedEventsInGroup(db, userId, groupId)
	if err != nil {
		fmt.Fprintf(w, "Error on getting event joined")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	events, err := GetEventsInGroup(db, userId, groupId)
	if err != nil {

		fmt.Fprintf(w, "Error on getting event")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	responseGroup := ResponseGroup{
		GroupData:   groupData,
		Posts:       allPosts,
		Members:     members,
		Followers:   membersFilt,
		Event:       events,
		EventJoined: eventJoined,
	}
	jsonData, err := json.Marshal(responseGroup)
	if err != nil {
		fmt.Fprintf(w, "Error encoding data to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
func CheckIfFollowerExists(db *sql.DB, groupID int, userID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM group_followers WHERE group_id = ? AND user_id = ?)`

	var exists bool
	err := db.QueryRow(query, groupID, userID).Scan(&exists)
	if err != nil {
		log.Printf("Erreur lors de la vérification de l'existence de l'utilisateur dans group_followers: %v", err)
		return false, err
	}

	return exists, nil
}

func GetEventsInGroup(db *sql.DB, userID, groupID int) ([]models.Event, error) {
	query := `
	SELECT event.*
		FROM event
		WHERE event.id NOT IN (
			SELECT event_id FROM event_joined WHERE user_id = ? AND group_id = ?
			) AND event.id NOT IN (
				SELECT event_id FROM event_notjoined WHERE user_id = ? AND group_id = ?
				) AND event.group_id = ?
	`

	rows, err := db.Query(query, userID, groupID, userID, groupID, groupID)
	if err != nil {
		log.Println("Erreur lors de l'exécution de la requête:", err)
		return nil, err
	}
	defer rows.Close()

	var neutralEvents []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.GroupID, &event.UserID, &event.Description, &event.Date, &event.Is_joined)
		if err != nil {
			log.Println("Erreur lors de la lecture des résultats:", err)
			return nil, err
		}
		neutralEvents = append(neutralEvents, event)
	}

	if err := rows.Err(); err != nil {
		log.Println("Erreur lors de la lecture des résultats:", err)
		return nil, err
	}

	return neutralEvents, nil
}

func GetJoinedEventsInGroup(db *sql.DB, userID, groupID int) ([]models.Event, error) {
	query := `
		SELECT event.*
		FROM event
		JOIN event_joined ON event.id = event_joined.event_id
		WHERE event_joined.user_id = ? AND event.group_id = ?
	`

	rows, err := db.Query(query, userID, groupID)
	if err != nil {
		log.Println("Erreur lors de l'exécution de la requête:", err)
		return nil, err
	}
	defer rows.Close()

	var joinedEvents []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.GroupID, &event.UserID, &event.Description, &event.Date, &event.Is_joined)
		if err != nil {
			fmt.Println(err)
			log.Println("Erreur lors de la lecture des résultats:", err)
			return nil, err
		}
		joinedEvents = append(joinedEvents, event)
	}

	if err := rows.Err(); err != nil {
		log.Println("Erreur lors de la lecture des résultats:", err)
		return nil, err
	}

	return joinedEvents, nil
}

func GetGroupIDFromRequest(w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.URL.Query().Get("id")
	if id == "" {
		fmt.Fprintf(w, "No id in URL")
		http.Error(w, "no ID", http.StatusBadRequest)
		return 0, fmt.Errorf("no group ID")
	}

	groupID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Invalid Id")
		return 0, fmt.Errorf("invalid group ID")
	}

	return groupID, nil
}

func GetUserData(db *sql.DB, userID int) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id =?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return user, fmt.Errorf("err execute user query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &user.Avatar, &user.About_me)
		if err != nil {
			return user, fmt.Errorf("err scan group data: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return user, fmt.Errorf("err read all groups results: %w", err)
	}
	return user, nil
}

func contains(arr []int, value int) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}
	return false
}

func getInfoGroup(db *sql.DB, groupId int) (models.Group, error) {
	var group models.Group
	// db.QueryRow("SELECT name, description, id_user_create FROM groups WHERE id = ?", groupId).Scan(&name, &id_user_create)

	query := "SELECT * FROM groups WHERE id = ?"
	rows, err := db.Query(query, groupId)
	if err != nil {
		return models.Group{}, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			return models.Group{}, fmt.Errorf("err scan group data: %w", err)
		}
	}

	group.User, err = GetUserData(db, group.ID_User_Create)
	if err != nil {
		return models.Group{}, fmt.Errorf("err get userData Group: %w", err)
	}
	return group, nil
}

func getMembers(db *sql.DB, groupID int) ([]models.User, error) {
	var members []models.User

	query := "SELECT user_id FROM group_followers WHERE group_id = ?"
	rows, err := db.Query(query, groupID)
	if err != nil {
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}
		user, err := GetUserData(db, id)
		if err != nil {
			return nil, fmt.Errorf("err execute all groups query: %w", err)
		}
		members = append(members, user)
	}

	return members, nil
}

func GetFollowers(db *sql.DB, userID int) ([]models.User, error) {
	var followers []models.User

	query := "SELECT user_id_follower FROM users_followers WHERE user_id_followed = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}
		user, err := GetUserData(db, id)
		if err != nil {
			return nil, fmt.Errorf("err execute all groups query: %w", err)
		}
		followers = append(followers, user)
	}

	return followers, nil
}

func formatDateTimeFr(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}

	weekday := t.Weekday().String()[:3]
	day := t.Day()
	month := (t.Month().String()[:3])
	year := t.Year()
	hour := t.Hour()

	return fmt.Sprintf("%s %d %s %d at %dH", weekday, day, month, year, hour)
}

func filterNonMembers(followers []models.User, members []models.User) []models.User {
	memberMap := make(map[int]bool)
	var nonMembers []models.User

	// Indexer les membres par ID dans une map pour la recherche rapide
	for _, member := range members {
		memberMap[member.Id] = true
	}

	// Vérifier chaque follower pour voir s'il n'est pas un membre
	for _, follower := range followers {
		if _, found := memberMap[follower.Id]; !found {
			nonMembers = append(nonMembers, follower)
		}
	}

	return nonMembers
}
