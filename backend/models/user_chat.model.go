package models

type Chat struct {
	ID              int    `json:"id"`
	UserSender      int    `json:"user_id_sender"`
	UserReceiver    *int   `json:"user_id_receiver"`
	MessageContent  string `json:"message_content"`
	GroupIDReceiver *int   `json:"group_id_receiver"`
	Date            string `json:"date"`
	Type            string `json:"type"`
	FirstName       string `json:"first_name"`
}

// type Message struct {
// 	UserIDSender   int    `json:"user_id_sender"`
// 	UserIDReceiver int    `json:"user_id_receiver"`
// 	MessageContent string `json:"message_content"`
// 	Date           string `json:"date"`
// }
