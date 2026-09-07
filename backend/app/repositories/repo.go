package repositories

import (
	db "backend/database"
)

type BaseRepo struct {
	DB        *db.Database
	TableName string
}

var (
	// UserRepo is the repository for user
	UserRepo = &UserRepository{}
	// SessionRepo is the repository for session
	SessionRepo = &SessionRepository{}
	PostRepo    = &PostRepository{}
	CommentRepo = &CommentRepository{}
	FollowRepo  = &FollowRepository{}
)

func init() {
	UserRepo.init()
	SessionRepo.init()
	PostRepo.init()
	CommentRepo.init()
	FollowRepo.init()
}
