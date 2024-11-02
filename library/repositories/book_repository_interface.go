package repositories

import "library/models"

type BookRepositoryInterface interface {
	Save(book *models.Book)
	FindByID(id string) (*models.Book, bool)
	FindAll() []*models.Book
	RemoveBook(bookID string)
}
