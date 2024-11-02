package repositories

import (
	"library/models"
	"sync"
)

type InMemoryLibrarianRepository struct {
	librarians map[string]*models.Librarian
	mtx        sync.RWMutex
}

func (lr *InMemoryLibrarianRepository) Save(librarian *models.Librarian) {
	lr.mtx.Lock()
	defer lr.mtx.Unlock()
	lr.librarians[librarian.ID] = librarian
}

func (lr *InMemoryLibrarianRepository) FindByID(id string) (*models.Librarian, bool) {
	lr.mtx.RLock()
	defer lr.mtx.RUnlock()
	librarian, ok := lr.librarians[id]
	return librarian, ok
}

func NewInMemoryLibrarianRepository() *InMemoryLibrarianRepository {
	return &InMemoryLibrarianRepository{
		librarians: make(map[string]*models.Librarian),
	}
}
