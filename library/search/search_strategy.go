package search

import "library/models"

type SearchStrategy interface {
	Search(criteria string, books []*models.Book) []*models.Book
}
