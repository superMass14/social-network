package groups

import (
	"backend/database"
	"backend/models"
	"backend/app/cors"
	"backend/app/service"
	utils "backend/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// func (w http.ResponseWriter, r *http.Request) {}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	// userID := 1
	cors.SetCors(&w)

	switch r.Method {
	case "POST":
		description := r.FormValue("content")
		dates := r.FormValue("date")
		heure := r.FormValue("heure")

		descrip := len(strings.TrimSpace(description))
		dat := len(strings.TrimSpace(dates))
		heur := len(strings.TrimSpace(heure))

		if descrip == 0 || dat == 0 || heur == 0 {
			log.Println("all fill required")
			return
		}

		date, err := formatDateTime(dates, heure)
		if err != nil {
			log.Println("Error on formatDateTime", err)
			return
		}

		session, err := service.AuthServ.VerifyToken(r)
		if err != nil {
			log.Println("Invalid Token ", err)
			utils.Alert(w, models.Errormessage{
				Type:       "Create Event",
				Msg:        "Invalid Token",
				StatusCode: http.StatusBadRequest,
			})
			return
		}

		userID := session.User_id

		groupID, err := GetGroupIDFromRequest(w, r)
		if err != nil {
			log.Println("Error on GetGroupIDFromRequest")
			return
		}

		_, err1 := insertEvent(groupID, userID, description, date)
		if err1 != nil {
			log.Println("problem after create service ", err)
			utils.Alert(w, models.Errormessage{
				Type:       "Insert event",
				Msg:        "Cannot insert event",
				StatusCode: http.StatusInternalServerError,
			})
			return
		} else {
			msg := models.Errormessage{
				Type:       "success",
				Msg:        "post created successfully",
				StatusCode: 200,
				Display:    false,
			}
			utils.Alert(w, msg)
		}

	}
}

func insertEvent(groupID, userID int, description string, eventDate string) (int64, error) {
	db := database.DB

	stmt, err := db.Prepare(`
    INSERT INTO event (group_id,user_id, description, event_date)
    VALUES (?, ?,?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(groupID, userID, description, eventDate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	stmt, err = db.Prepare(`
    INSERT INTO event_joined (event_id, user_id, group_id)
    VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, userID, groupID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func formatDateTime(date, hour string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %w", err)
	}

	// Ensure hour is a valid integer between 0 and 23
	var parsedHour int
	if _, err := fmt.Sscan(hour, &parsedHour); err != nil || parsedHour < 0 || parsedHour > 23 {
		return "", fmt.Errorf("invalid hour format: %w", err)
	}

	// Combine date and hour into a single timestamp
	targetTime := parsedDate.Add(time.Hour * time.Duration(parsedHour))

	// Format the timestamp for SQLite3 (YYYY-MM-DD HH:MM:SS)
	return targetTime.Format("2006-01-02 15:04:05"), nil
}
