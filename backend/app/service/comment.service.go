package service

import (
	"backend/models"
	r "backend/app/repositories"
)

type CommentService struct {
	CommentRepo r.CommentRepository
}

func (c *CommentService) init() {
	c.CommentRepo = *r.CommentRepo
}

func (c *CommentService) CreateComment(comment models.Comment) (bool, models.Errormessage) {
	Not_ok, err := c.CommentRepo.CreateComment(comment)
	if Not_ok {
		return true, err
	}
	return false, models.Errormessage{}
}

func (c *CommentService) FetchComments(Id string) ([]models.DataComment, error) {
	comments, err := c.CommentRepo.LoadComment(Id)
	if comments == nil {
		return nil, err
	}
	return comments, nil
}
