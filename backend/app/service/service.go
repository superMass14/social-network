package service

var (
	AuthServ    AuthService
	PostServ    PostService
	CommentServ CommentService
)

func init() {
	AuthServ.init()
	PostServ.init()
	CommentServ.init()
}
