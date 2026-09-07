package models

// type Group struct {
// 	ID           int    `json:"id"`
// 	Name         string `json:"group_name"`
// 	Description  string `json:"group_description"`
// 	UserIDOwner  int    `json:"user_id_owner"`
// 	DateCreation string `json:"dateCreation"`
// 	Avatar       string `json:"avatar"`
// }

type Group struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ID_User_Create int    `json:"id_user_create"`
	User           User   `json:"creator"`
	Avatar         string `json:"image"`
	Creation_Date  string `json:"creation_date"`
	Href           string `json:"href"`
	State           string `json:"state"`
}
