package repositories

import (
	"library/models"
	"sync"
)

type InMemoryBookRepository struct {
	books map[string]*models.Book
	mu    sync.RWMutex
}

func (br *InMemoryBookRepository) Save(book *models.Book) {
	br.mu.Lock()
	defer br.mu.Unlock()

	br.books[book.ID] = book
}

func (br *InMemoryBookRepository) RemoveBook(bookID string) {
	delete(br.books, bookID)
}

func (br *InMemoryBookRepository) FindByID(id string) (*models.Book, bool) {
	br.mu.RLock()
	defer br.mu.RUnlock()
	book, ok := br.books[id]
	return book, ok
}

func (br *InMemoryBookRepository) FindAll() []*models.Book {
	br.mu.RLock()
	defer br.mu.RUnlock()

	books := make([]*models.Book, 0)
	for _, book := range br.books {
		books = append(books, book)
	}

	return books
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{
		books: make(map[string]*models.Book),
	}
}
