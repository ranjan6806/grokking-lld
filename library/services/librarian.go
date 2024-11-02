package services

import (
	"library/models"
	"library/repositories"
)

type LibrarianService struct {
	BookRepo   repositories.BookRepositoryInterface
	MemberRepo repositories.MemberRepositoryInterface
}

func (lbs *LibrarianService) AddBook(book *models.Book) {
	lbs.BookRepo.Save(book)
}

func (lbs *LibrarianService) RemoveBook(bookID string) {
	lbs.BookRepo.RemoveBook(bookID)
}

func (lbs *LibrarianService) AddMember(member *models.Member) {
	lbs.MemberRepo.Save(member)
}

func (lbs *LibrarianService) RemoveMember(memberID string) {
	lbs.MemberRepo.RemoveMember(memberID)
}

func NewLibrarianService(bookRepo repositories.BookRepositoryInterface, memberRepo repositories.MemberRepositoryInterface) LibrarianServiceInterface {
	return &LibrarianService{
		BookRepo:   bookRepo,
		MemberRepo: memberRepo,
	}
}
