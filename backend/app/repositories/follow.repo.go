package repositories

import (
	db "backend/database"
	opt "backend/database/operators"
	q "backend/database/query"
	"backend/models"
	"fmt"
)

type FollowRepository struct {
	BaseRepo
}

func (f *FollowRepository) init() {
	f.DB = db.DB
	f.TableName = "users_followers"
}

func (f *FollowRepository) Getfollower(id int) ([]models.UserFollower, error) {
	rows, err := f.DB.GetAllFrom(f.TableName, q.WhereOption{"user_id_followed": opt.Equals(id)}, "", nil)
	if err != nil {
		return nil, err
	}
	var IdUsers []models.UserFollower
	for rows.Next() {
		var iduser models.UserFollower
		err := rows.Scan(&iduser.User_id_followed, &iduser.User_id_follower)
		if err != nil {
			return nil, err
		}
		IdUsers = append(IdUsers, iduser)
	}
	return IdUsers, nil
}

func (f *FollowRepository) Getfollowing(id int) ([]models.UserFollower, error) {
	rows, err := f.DB.GetAllFrom(f.TableName, q.WhereOption{"user_id_follower": opt.Equals(id)}, "", nil)
	if err != nil {
		return nil, err
	}
	var IdUsers []models.UserFollower
	for rows.Next() {
		var iduser models.UserFollower
		err := rows.Scan(&iduser.User_id_followed, &iduser.User_id_follower)
		if err != nil {
			return nil, err
		}
		IdUsers = append(IdUsers, iduser)
	}
	return IdUsers, nil
}

func (f *FollowRepository) FollowUser(id int, idToFollow int) error {
	verif, err := VerifyUserFollow(db.DB, idToFollow, id)
	if err != nil {
		fmt.Println("Erreur selected users_followers ", err)
	}
	if !verif {
		err := f.DB.Insert(f.TableName, models.UserFollower{User_id_followed: idToFollow, User_id_follower: id})
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FollowRepository) UnfollowUser(followedID, followerID int) error {
	err := f.DB.Delete(f.TableName, q.WhereOption{"user_id_followed": opt.Equals(followedID), "user_id_follower": opt.Equals(followerID)})
	if err != nil {
		return err
	}
	return nil
}

func VerifyUserFollow(DB *db.Database, user_id_followed, user_id_follower int) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users_followers WHERE user_id_followed = $1 AND user_id_follower = $2", user_id_followed, user_id_follower).Scan(&count)
	if err != nil {
		fmt.Println("Erreur requete ", err)
		return false, err
	}
	return count > 0, nil
}
