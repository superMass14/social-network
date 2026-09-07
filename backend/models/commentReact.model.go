package models

type CommentReact struct {
	ID int
	Post_id    int  
	Comment_id int  
	User_id    int  
	Reaction   bool
}
