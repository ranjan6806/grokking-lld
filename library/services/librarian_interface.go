package services

import "library/models"

type LibrarianServiceInterface interface {
	AddBook(book *models.Book)
	RemoveBook(bookID string)
	AddMember(member *models.Member)
	RemoveMember(memberID string)
}
