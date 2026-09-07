package repositories

import (
	db "backend/database"
	opt "backend/database/operators"
	q "backend/database/query"
	"backend/models"
	"database/sql"
	"fmt"
	"log"
)

type UserRepository struct {
	BaseRepo
}

func (u *UserRepository) init() {
	u.DB = db.DB
	u.TableName = "users"
}

func (u *UserRepository) SaveUser(user models.FormatedUser) error {
	err := u.DB.Insert(u.TableName, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserById(id int) (models.User, error) {
	var user models.User
	row, err := u.DB.GetOneFrom(u.TableName, q.WhereOption{"id": opt.Equals(id)})
	if err == sql.ErrNoRows {
		return models.User{}, err
	}

	var user_name sql.NullString
	var gender sql.NullString
	var avatar sql.NullString
	var about_me sql.NullString
	err = row.Scan(&user.Id, &user.First_name, &user.Last_name, &user_name, &gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &avatar, &about_me)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
		log.Println("Error scanning row", err)
		return models.User{}, err
	}
	user.User_name = GetStringValue(user_name)
	user.Gender = GetStringValue(gender)
	user.Avatar = GetStringValue(avatar)
	user.About_me = GetStringValue(about_me)
	return user, nil
}
func (u *UserRepository) GetUserByToken(token string) (models.User, error) {
	var user models.User
	row, err := u.DB.GetOneFrom(u.TableName, q.WhereOption{"token": opt.Equals(token)})
	if err == sql.ErrNoRows {
		log.Println("No rows found 0", err)
		return models.User{}, err
	}
	var user_name sql.NullString
	var gender sql.NullString
	var avatar sql.NullString
	var about_me sql.NullString
	err = row.Scan(&user.Id, &user.First_name, &user.Last_name, &user_name, &gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &avatar, &about_me)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found", err)
			return models.User{}, err
		}
		log.Println("Error scanning row", err)
		return models.User{}, err
	}
	user.User_name = GetStringValue(user_name)
	user.Gender = GetStringValue(gender)
	user.Avatar = GetStringValue(avatar)
	user.About_me = GetStringValue(about_me)

	return user, nil
}

func (u *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row, err := u.DB.GetOneFrom(u.TableName, q.WhereOption{"email": opt.Equals(email)})
	if err != nil {
		log.Println("Error getting user by email:", err)
		return models.User{}, fmt.Errorf("error getting user by email: %v", err)
	}
	var user_name sql.NullString
	var gender sql.NullString
	var avatar sql.NullString
	var about_me sql.NullString
	err = row.Scan(&user.Id, &user.First_name, &user.Last_name, &user_name, &gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &avatar, &about_me)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found", err)
			return models.User{}, err
		}
		log.Println("Error scanning row", err)
		return models.User{}, err
	}
	user.User_name = GetStringValue(user_name)
	user.Gender = GetStringValue(gender)
	user.Avatar = GetStringValue(avatar)
	user.About_me = GetStringValue(about_me)

	return user, nil
}

func (u *UserRepository) UpdateUser(user models.User) error {
	err := u.DB.Update(u.TableName, user, q.WhereOption{"id": opt.Equals(user.Id)})
	return err
}

func (u *UserRepository) DeleteUser(user models.User) error {
	err := u.DB.Delete(u.TableName, q.WhereOption{"id": opt.Equals(user.Id)})
	return err
}

func (u *UserRepository) GetAllUsers() (users []models.User, err error) {
	var user models.User
	rows, err := u.DB.GetAllFrom(u.TableName, nil, "email", nil)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user_name sql.NullString
		var gender sql.NullString
		var avatar sql.NullString
		var about_me sql.NullString
		err = rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user_name, &gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &avatar, &about_me)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No rows found", err)
				return nil, err
			}
			log.Println("Error scanning row", err)
			return nil, err
		}
		user.User_name = GetStringValue(user_name)
		user.Gender = GetStringValue(gender)
		user.Avatar = GetStringValue(avatar)
		user.About_me = GetStringValue(about_me)
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepository) UpdateProfileType(user models.User) error {
	err := u.DB.Update(u.TableName, user, q.WhereOption{"id": opt.Equals(user.Id)})
	return err
}

func GetStringValue(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}
