package models

type Notification struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	Status           bool   `json:"status"`
	UserIDSender     int    `json:"user_id_sender"`
	UserIDReceveived int    `json:"user_id_received"`
	GroupId          int    `json:"group_id"`
	Date             string `json:"date"`
}

type NotificationUser struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	Status           bool   `json:"status"`
	UserIDSender     int    `json:"user_id_sender"`
	UserIDReceveived int    `json:"user_id_received"`
	GroupId          int    `json:"group_id"`
	Date             string `json:"date"`
	First_name       string `json:"first_name"`
	Last_name        string `json:"last_name"`
	User_name        string `json:"user_name"`
	Email            string `json:"email"`
}
