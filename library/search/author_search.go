package search

import "library/models"

type AuthorSearch struct{}

func (s *AuthorSearch) Search(criteria string, books []*models.Book) []*models.Book {
	var results []*models.Book
	for _, book := range books {
		if book.Author == criteria {
			results = append(results, book)
		}
	}

	return results
}
