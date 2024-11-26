package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	filebasedDb *filebased.Data
}

func NewSessionsRepo(filebasedDb *filebased.Data) *sessionsRepo {
	return &sessionsRepo{filebasedDb}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	err := u.filebasedDb.AddSession(session)
	if err != nil {
		return fmt.Errorf("failed to add session: %v", err)
	}
	return nil
}

func (u *sessionsRepo) DeleteSession(token string) error {
	err := u.filebasedDb.DeleteSession(token)
	if err != nil {
		return fmt.Errorf("failed to delete session: %v", err)
	}

	return nil
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	err := u.filebasedDb.UpdateSession(session)
	if err != nil {
		return fmt.Errorf("failed to update session: %v", err)
	}

	return nil
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	session, err := u.filebasedDb.SessionAvailEmail(email)
	if err != nil {
		return model.Session{}, fmt.Errorf("session not found for email %s: %v", email, err)
	}
	return session, nil
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	session, err := u.filebasedDb.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, fmt.Errorf("session not found for token %s: %v", token, err)
	}
	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
