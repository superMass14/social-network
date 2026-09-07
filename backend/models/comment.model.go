package models

type Comment struct {
	ID      string
	Date    string
	Content string
	Post_id string
	User_id int
	Media   string
}

type DataComment struct {
	Avatar   string
	Content  string
	Date     string
	FullName string
	UserName string
	ID       int
	Media    string
	Post_id  string
	Username string
}
