package services

import "library/models"

type LibraryServiceInterface interface {
	BorrowBook(bookID, userID string)
	ReturnBook(bookCopyID, userID string)
	SearchBooks(criteria string) []*models.Book
}
