package models

import "time"

type Session struct {
	Token      string    `json:"token"`
	User_id    int       `json:"user_id"`
	Expiration time.Time `json:"expiration"`
}
