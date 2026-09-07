package models

type WResponse struct {
	Type       string
	Data       any
	Followers  []User
	StatusCode int
	Display    bool
	Msg        string
}
