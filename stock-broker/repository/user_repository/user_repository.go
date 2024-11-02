package user_repository

import (
	"errors"
	"stock-broker/internal/user"
)

type UserRepository interface {
	AddUser(user *user.User) error
	GetUser(userID string) (*user.User, error)
}

type InMemoryUserRepository struct {
	users map[string]*user.User
}

func (ur *InMemoryUserRepository) AddUser(user *user.User) error {
	if _, exists := ur.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	ur.users[user.ID] = user
	return nil
}

func (ur *InMemoryUserRepository) GetUser(userID string) (*user.User, error) {
	user, exists := ur.users[userID]
	if !exists {
		return nil, errors.New("user does not exists")
	}

	return user, nil
}

func NewInMemoryUserRepository() UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*user.User),
	}
}
