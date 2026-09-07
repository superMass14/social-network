package models

type Confirm struct {
	UserIDAsker int `json:"user_id"`
	UserIDAsked int `json:"user_id_answerer"`
	GroupId     int `json:"groupId"`
}
