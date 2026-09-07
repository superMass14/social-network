package models

type UserFollower struct {
	User_id_followed int `json:"user_id_followed"`
	User_id_follower int `json:"user_id_follower"`
}
