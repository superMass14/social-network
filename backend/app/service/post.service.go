package service

import (
	"backend/models"
	r "backend/app/repositories"
	"log"
)

type PostService struct {
	PostRepo r.PostRepository
}

func (p *PostService) init() {
	p.PostRepo = *r.PostRepo
}

func (p *PostService) CreatePost(post models.Post) (bool, models.Errormessage) {
	log.Println("creating with => ", post.ToIns.User_id)
	Not_ok, err := p.PostRepo.CreatePost(post)
	if Not_ok {
		return true, err
	}
	return false, models.Errormessage{}
}

func (p *PostService) GetPost(Id int) ([]models.DataPost, error) {
	log.Println("getting post with ", Id)
	posts, err := p.PostRepo.LoadPost(Id)
	if posts == nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostService) FetchPostGroup(Id int) ([]models.DataPost, error) {
	post, err := p.PostRepo.LoadPostGroup(Id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (p *PostService) FetchPostGroupByUserID(CurrentId int, idProfil int) ([]models.DataPost, error) {
	post, err := p.PostRepo.LoadPostGroupByUserID(CurrentId, idProfil)
	if err != nil {
		return nil, err
	}
	return post, nil
}
