package models

type Event struct {
	ID          int    `json:"id"`
	Event_id    int    `json:"event_id"`
	GroupID     int    `json:"group_id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Is_joined   int    `json:"is_joined"`
}

// type EventJoined struct {
// 	ID          int    `json:"id"`
// 	User_id     int    `json:"user"`
// 	Group_id    int    `json:"group"`
// 	Description string `json:"description"`
// 	Date        string `json:"date"`
// }
