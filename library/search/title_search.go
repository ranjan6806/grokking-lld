package search

import "library/models"

type TitleSearch struct {
}

func (s *TitleSearch) Search(criteria string, books []*models.Book) []*models.Book {
	var results []*models.Book
	for _, book := range books {
		if book.Title == criteria {
			results = append(results, book)
		}
	}
	return results
}
