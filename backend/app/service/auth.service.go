package service

import (
	"backend/models"
	r "backend/app/repositories"
	"backend/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo   r.UserRepository
	SessRepo   r.SessionRepository
	FollowRepo r.FollowRepository
}

func (a *AuthService) init() {
	a.UserRepo = *r.UserRepo
	a.SessRepo = *r.SessionRepo
	a.FollowRepo = *r.FollowRepo
}

func (a *AuthService) CreateUser(user models.FormatedUser) error {
	_, err := a.UserRepo.GetUserByEmail(user.Email)
	if err == nil {
		return fmt.Errorf("this email is already in use")
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	err = a.UserRepo.SaveUser(user)
	return err
}

func (a *AuthService) CheckCredentials(email, password string) (models.User, error) {
	user, err := a.UserRepo.GetUserByEmail(email)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid Email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, fmt.Errorf("invalid Password")
	}
	return user, nil
}

func (a *AuthService) VerifyTokenStr(token string) (userID int, err error) {
	if token == "" {
		return 0, fmt.Errorf("missing token")
	}

	session, err := a.SessRepo.GetSession(token)
	if err != nil {
		log.Println("Error getting session", err)
		return 0, err
	}

	if session.Expiration.Before(time.Now()) {
		log.Println("token expired")
		return 0, fmt.Errorf("token expired")
	}

	return session.User_id, nil
}

func (a *AuthService) VerifyToken(r *http.Request) (session models.Session, err error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		token, err = url.QueryUnescape(r.URL.Query().Get("token"))
		if err != nil {
			log.Println("Erreur token", err.Error())
			return models.Session{}, fmt.Errorf("invalid token")
		}
		token = strings.ReplaceAll(token, " ", "+")
	}

	if token == "" {
		return models.Session{}, fmt.Errorf("missing token")
	}
	session, err = a.SessRepo.GetSession(token)
	if err != nil {
		log.Println("Error getting session", err)
		return models.Session{}, err
	}
	if session.Expiration.Before(time.Now()) {
		log.Println("token expired")
		return models.Session{}, fmt.Errorf("token expired")
	}
	return session, nil
}

func (a *AuthService) RemoveSession(token string) error {
	return a.SessRepo.DeleteSession(token)
}

func (a *AuthService) RemExistingSession(userId int) error {
	sessions, err := a.SessRepo.GetSessionByUserId(userId)
	if err != nil {
		return err
	}

	for _, session := range sessions {
		err = a.SessRepo.DeleteSession(session.Token)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *AuthService) CreateSession(user models.User) (models.Session, error) {
	err := a.RemExistingSession(user.Id)
	if err != nil {
		log.Println("Error removing existing session", err)
		return models.Session{}, err
	}
	token := utils.GenerateToken()
	expiration := time.Now().Add(3 * time.Hour)
	err = a.SessRepo.SaveSession(models.Session{Token: token, User_id: user.Id, Expiration: expiration})
	if err != nil {
		log.Println("Error saving session", err)
		return models.Session{}, err
	}
	return models.Session{Token: token, User_id: user.Id, Expiration: expiration}, nil

}

func (a *AuthService) GetUserById(id int) (models.User, error) {
	user, err := a.UserRepo.GetUserById(id)
	if err != nil {
		log.Println("Error getting user by id", err)
		return models.User{}, err
	}
	return user, nil
}

func (a *AuthService) GetFollowersByID(id int) ([]models.User, error) {
	IdUsers, err := a.FollowRepo.Getfollower(id)
	if err != nil {
		return nil, err
	}
	followers := make([]models.User, 0)
	for _, idUser := range IdUsers {
		user, err := a.UserRepo.GetUserById(idUser.User_id_follower)
		if err != nil {
			return nil, err
		}
		followers = append(followers, user)
	}
	return followers, nil
}

func (a *AuthService) GetFollowingsByID(id int) ([]models.User, error) {
	IdUsers, err := a.FollowRepo.Getfollowing(id)
	if err != nil {
		return nil, err
	}
	followings := make([]models.User, 0)
	for _, idUser := range IdUsers {
		user, err := a.UserRepo.GetUserById(idUser.User_id_followed)
		if err != nil {
			return nil, err
		}
		followings = append(followings, user)
	}
	return followings, nil
}

func (a *AuthService) FollowUser(followerID, followingID int) error {
	err := a.FollowRepo.FollowUser(followerID, followingID)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) UnfollowUser(followedID, followerID int) error {
	err := a.FollowRepo.UnfollowUser(followedID, followerID)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthService) UpdateProfileType(id int, profileType string) error {
	user, err := a.GetUserById(id)
	if err != nil {
		return err
	}
	if profileType == "public" {
		user.User_type = "private"
	} else {
		user.User_type = "public"
	}
	err = a.UserRepo.UpdateProfileType(user)
	if err != nil {
		return err
	}
	return nil
}
