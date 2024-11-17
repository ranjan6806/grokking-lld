package search

import "movie-ticket-booking/models"

type SearchStrategy interface {
	Search(cinema *models.Cinema) []*models.Movie
}

type TitleSearchStrategy struct {
}

func (ts *TitleSearchStrategy) Search(cinema *models.Cinema) []*models.Show {
	return nil
}

type GenreSearchStrategy struct{}

func (gs *GenreSearchStrategy) Search(cinema *models.Cinema) []*models.Show {
	return nil
}

type LanguageSearchStrategy struct{}

func (ls *LanguageSearchStrategy) Search(cinema *models.Cinema) []*models.Show {
	return nil
}
