package handler

import (
	"backend/models"
	"backend/app/cors"
	"backend/app/service"
	"backend/utils"
	"backend/utils/seed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetUserGroup(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodGet:
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
		var tableauVide [][]models.User
		listGroup, _ := seed.GetGroup(seed.DB, session.User_id)
		listUser, err := seed.ListeUsers(seed.DB, session.User_id)
		if err != nil {
			fmt.Println("error:", err)
		} else if len(listUser) == 0 && len(listGroup) == 0 {
			fmt.Println("La liste des utilisateurs est vide")
			data := []interface{}{tableauVide, tableauVide}
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Fprintf(w, "Error encoding data to JSON: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else if len(listUser) != 0 && len(listGroup) == 0 {
			data := []interface{}{listUser[0], tableauVide}
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Fprintf(w, "Error encoding data to JSON: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else if len(listUser) == 0 && len(listGroup) != 0 {
			data := []interface{}{tableauVide, listGroup}
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Fprintf(w, "Error encoding data to JSON: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		} else {
			data := []interface{}{listUser[0], listGroup}
			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Fprintf(w, "Error encoding data to JSON: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}

	}
}
