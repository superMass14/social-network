package models

type Post struct {
	ToIns      InsPost
	Categories []string
	Viewers    string
}
type InsPost struct {
	ID       string
	Date     string
	Content  string
	Media    string
	User_id  int
	Group_id int
	Privacy  string
}

type Category struct {
	Post_id  string 
	Category string
}

type DataPost struct {
	Avatar     string
	Categories string
	Comments   int
	Content    string
	Date       string
	FullName   string
	Group_id   int
	ID         string
	Media      string
	Privacy    string
	UserName   string
	User_id    int
	Viewers    string
}
