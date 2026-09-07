package notification

import (
	"backend/app/cors"
	joingroup "backend/app/handler/groups/joinGroup"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetNotification(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// db := database.NewDatabase()
	listeNotificaton, err := joingroup.NotificationUserData(r)
	if err != nil {
		fmt.Println("Erreur check notification ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listeNotificaton)
}
