package repositories

import "library/models"

type BookCopyRepositoryInterface interface {
	Save(copy *models.BookCopy)
	FindByID(id string) (*models.BookCopy, bool)
	FindAvailableCopy(bookID string) (*models.BookCopy, bool)
	DecreaseAvailableCopies(bookID string)
	IncreaseAvailableCopies(bookID string)
}
