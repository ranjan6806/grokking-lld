package models

import "time"

type Movie struct {
	Title       string
	Language    string
	Genre       string
	ReleaseDate time.Time
}

func NewMovie(title, language, genre string, releaseDate time.Time) *Movie {
	return &Movie{
		Title:       title,
		Language:    language,
		Genre:       genre,
		ReleaseDate: releaseDate,
	}
}
