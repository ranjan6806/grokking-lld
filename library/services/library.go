package services

import (
	"library/models"
	"library/repositories"
	"library/search"
)

type LibraryService struct {
	BookRepo       repositories.BookRepositoryInterface
	BookCopyRepo   repositories.BookCopyRepositoryInterface
	MemberRepo     repositories.MemberRepositoryInterface
	LibrarianRepo  repositories.LibrarianRepositoryInterface
	SearchStrategy search.SearchStrategy
}

func NewLibraryService(
	bookRepo repositories.BookRepositoryInterface,
	bookCopyRepo repositories.BookCopyRepositoryInterface,
	memberRepo repositories.MemberRepositoryInterface,
	librarianRepo repositories.LibrarianRepositoryInterface,
	searchStrategy search.SearchStrategy,
) LibraryServiceInterface {
	return &LibraryService{
		BookRepo:       bookRepo,
		BookCopyRepo:   bookCopyRepo,
		MemberRepo:     memberRepo,
		LibrarianRepo:  librarianRepo,
		SearchStrategy: searchStrategy,
	}
}

func (s *LibraryService) BorrowBook(bookID, userID string) {
	member, memberExists := s.MemberRepo.FindByID(userID)
	if !memberExists {
		return
	}

	bookCopy, available := s.BookCopyRepo.FindAvailableCopy(bookID)
	if !available {
		return
	}

	if bookCopy.Issue(member) {
		s.MemberRepo.IssueBook(userID, bookCopy)
	}
}

func (s *LibraryService) ReturnBook(bookID, userID string) {
	bookCopy, exists := s.BookCopyRepo.FindByID(bookID)
	if !exists {
		return
	}

	member, memberExists := s.MemberRepo.FindByID(userID)
	if !memberExists {
		return
	}

	bookCopy.Return()
	member.ReturnBook(bookID)
}

func (s *LibraryService) SearchBooks(criteria string) []*models.Book {
	allBooks := s.BookRepo.FindAll()
	return s.SearchStrategy.Search(criteria, allBooks)
}

func (s *LibraryService) borrowBookForMember(bookID string, member *models.Member) {
	if !member.CanIssueBook() {
		return
	}

	bookCopy, available := s.BookCopyRepo.FindAvailableCopy(bookID)
	if !available {
		return
	}

	if bookCopy.Issue(member) {
		member.IssueBook(bookCopy)
		return
	}
}
