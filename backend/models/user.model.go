package models

type User struct {
	Id        int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	User_name  string `json:"user_name"`
	Gender     string `json:"gender"`
	User_type  string `json:"user_type"`
	Birth_date string `json:"birth_date"`
	Avatar     string `json:"avatar"`
	About_me   string `json:"about_me"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}

type FormatedUser struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	User_name  string `json:"user_name"`
	Gender     string `json:"gender"`
	User_type  string `json:"user_type"`
	Birth_date string `json:"birth_date"`
	Avatar     string `json:"avatar"`
	About_me   string `json:"about_me"`
	Password   string `json:"password"`
	Email      string `json:"email"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}