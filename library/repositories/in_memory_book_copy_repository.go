package repositories

import (
	"library/models"
	"sync"
)

type InMemoryBookCopyRepository struct {
	copies map[string]*models.BookCopy
	mtx    sync.RWMutex
}

func (bcr *InMemoryBookCopyRepository) Save(copy *models.BookCopy) {
	bcr.mtx.Lock()
	defer bcr.mtx.Unlock()
	bcr.copies[copy.ID] = copy
}

func (bcr *InMemoryBookCopyRepository) FindByID(id string) (*models.BookCopy, bool) {
	bcr.mtx.RLock()
	defer bcr.mtx.RUnlock()
	bookCopy, ok := bcr.copies[id]
	return bookCopy, ok
}

func (bcr *InMemoryBookCopyRepository) FindAvailableCopy(bookID string) (*models.BookCopy, bool) {
	bcr.mtx.RLock()
	defer bcr.mtx.RUnlock()
	for _, bookCopy := range bcr.copies {
		if bookCopy.BookID == bookID && bookCopy.IsAvailable() {
			return bookCopy, true
		}
	}

	return nil, false
}

func (bcr *InMemoryBookCopyRepository) DecreaseAvailableCopies(bookID string) {
	bcr.mtx.Lock()
	defer bcr.mtx.Unlock()
	for _, bookCopy := range bcr.copies {
		if bookCopy.BookID == bookID && bookCopy.IsAvailable() {
			bookCopy.Status = models.Issued
		}
	}
}

func (bcr *InMemoryBookCopyRepository) IncreaseAvailableCopies(bookID string) {
	bcr.mtx.Lock()
	defer bcr.mtx.Unlock()
	for _, bookCopy := range bcr.copies {
		if bookCopy.BookID == bookID && bookCopy.Status == models.Issued {
			bookCopy.Status = models.Available
		}
	}
}

func NewInMemoryBookCopyRepository() *InMemoryBookCopyRepository {
	return &InMemoryBookCopyRepository{copies: make(map[string]*models.BookCopy)}
}
