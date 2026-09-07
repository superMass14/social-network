package utils

import (
	"backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func Alert(w http.ResponseWriter, msg models.Errormessage) {
	response := map[string]interface{}{}
	response["msg"] = msg.Msg
	response["status"] = msg.StatusCode
	response["type"] = msg.Type
	response["display"] = msg.Display
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(msg.StatusCode)
	log.Printf("[Alert]: %v, %v\n", response["msg"], response["status"])
	json.NewEncoder(w).Encode(response)
}

func AlertData(w http.ResponseWriter, msg models.WResponse) {
	response := map[string]interface{}{}
	response["msg"] = msg.Msg
	response["status"] = msg.StatusCode
	response["type"] = msg.Type
	response["display"] = msg.Display
	response["data"] = msg.Data
	response["followers"] = msg.Followers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(msg.StatusCode)
	//log.Println("[Alert Data sent]: ", response["data"])
	json.NewEncoder(w).Encode(response)
}
