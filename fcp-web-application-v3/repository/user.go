package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	filebasedDb *filebased.Data
}

func NewUserRepo(filebasedDb *filebased.Data) *userRepository {
	return &userRepository{filebasedDb}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	user, err := r.filebasedDb.GetUserByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to find user by email %s: %v", email, err)
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	createdUser, err := r.filebasedDb.CreateUser(user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to create user: %v", err)
	}
	return createdUser, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	userTaskCategories, err := r.filebasedDb.GetUserTaskCategories()
	if err != nil {
		return nil, fmt.Errorf("failed to get user task categories: %v", err)
	}
	return userTaskCategories, nil
}
