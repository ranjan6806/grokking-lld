package main

import (
	"library/models"
	"library/repositories"
	"library/search"
	"library/services"
)

func main() {
	// Initialize repositories
	bookRepo := repositories.NewInMemoryBookRepository()
	bookCopyRepo := repositories.NewInMemoryBookCopyRepository()
	memberRepo := repositories.NewInMemoryMemberRepository()
	librarianRepo := repositories.NewInMemoryLibrarianRepository()

	libraryService := services.NewLibraryService(
		bookRepo,
		bookCopyRepo,
		memberRepo,
		librarianRepo,
		&search.TitleSearch{},
	)

	librarianService := services.NewLibrarianService(bookRepo, memberRepo)

	newBook := &models.Book{ID: "1", Title: "Go Programming", Author: "Alice Doe", Subject: "Programming"}
	librarianService.AddBook(newBook)

	newMember := &models.Member{ID: "M1", Name: "John Smith", LibraryCard: "LC123"}
	librarianService.AddMember(newMember)

	libraryService.BorrowBook("1", "M1")

}
