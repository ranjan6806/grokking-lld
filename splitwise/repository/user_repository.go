package repository

import (
	"fmt"
	"splitwise/models"
)

type InMemoryUserRepository struct {
	Users       map[string]*models.User
	userCounter int
}

func (ur *InMemoryUserRepository) AddUser(user *models.User) {
	userID := fmt.Sprintf("u%d", ur.userCounter)
	user.ID = userID
	ur.Users[userID] = user
	ur.userCounter++
}

func (ur *InMemoryUserRepository) RemoveUser(userID string) {
	delete(ur.Users, userID)
}

func (ur *InMemoryUserRepository) GetUser(userID string) *models.User {
	return ur.Users[userID]
}

func NewInMemoryUserRepository() UserRepositoryInterface {
	return &InMemoryUserRepository{
		Users:       make(map[string]*models.User),
		userCounter: 1,
	}
}
