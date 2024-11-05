package service

import (
	"fmt"
	"url-shortener/encoder"
	"url-shortener/repository"
)

type URLShortenerService struct {
	encoder encoder.Encoder
	repo    repository.URLRepository
}

func NewURLShortenerSService(encoder encoder.Encoder, repo repository.URLRepository) *URLShortenerService {
	return &URLShortenerService{
		encoder: encoder,
		repo:    repo,
	}
}

func (s *URLShortenerService) ShortenURL(url string) (string, error) {
	shortURL := s.encoder.Encode(url)
	err := s.repo.Save(shortURL, url)
	if err != nil {
		fmt.Println("error saving url", err)
		return "", err
	}

	return shortURL, nil
}

func (s *URLShortenerService) Fetch(shortURL string) (string, error) {
	return s.repo.Find(shortURL)
}
