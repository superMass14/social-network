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

type SessionRepository struct {
	BaseRepo
}

func (s *SessionRepository) init() {
	s.DB = db.DB
	s.TableName = "sessions"
}

func (s *SessionRepository) SaveSession(session models.Session) error {
	err := s.DB.Insert(s.TableName, session)
	return err
}

func (s *SessionRepository) GetSession(token string) (models.Session, error) {
	var session models.Session
	row, err := s.DB.GetOneFrom(s.TableName, q.WhereOption{"token": opt.Equals(token)})
	if err != nil {
		return session, err
	}
	err = row.Scan(&session.Token, &session.User_id, &session.Expiration)
	if err != nil {
		if err == sql.ErrNoRows {
			return session, fmt.Errorf("no session found with this userId")
		}
		return session, err
	}
	return session, nil
}

func (s *SessionRepository) GetSessionByUserId(userId int) (sessions []models.Session, err error) {
	var session models.Session
	rows, err := s.DB.GetAllFrom(s.TableName, q.WhereOption{"user_id": opt.Equals(userId)}, "", nil)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&session.Token, &session.User_id, &session.Expiration)
		sessions = append(sessions, session)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no session found with this userId")
		}
		return nil, err
	}
	return sessions, nil
}

func (s *SessionRepository) DeleteSession(token string) error {
	err := s.DB.Delete(s.TableName, q.WhereOption{"token": opt.Equals(token)})
	if err != nil {
		log.Println("Delete session error", err)
		return err
	}
	return nil
}

func (s *SessionRepository) UpdateSession(session models.Session) error {
	err := s.DB.Update(s.TableName, session, q.WhereOption{"token": opt.Equals(session.Token)})
	if err != nil {
		return err
	}
	return nil
}
