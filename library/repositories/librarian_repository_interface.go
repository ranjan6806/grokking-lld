package repositories

import "library/models"

type LibrarianRepositoryInterface interface {
	Save(librarian *models.Librarian)
	FindByID(id string) (*models.Librarian, bool)
}
