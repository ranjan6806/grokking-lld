package repository

import "errors"

type InMemoryURLRepository struct {
	data map[string]string
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		data: make(map[string]string),
	}
}

func (r *InMemoryURLRepository) Save(shortURL, originalURL string) error {
	if _, exists := r.data[shortURL]; exists {
		return errors.New("short URL already exists")
	}

	r.data[shortURL] = originalURL
	return nil
}

func (r *InMemoryURLRepository) Find(shortURL string) (string, error) {
	if _, exists := r.data[shortURL]; exists {
		return r.data[shortURL], nil
	}

	return "", errors.New("short URL not found")
}
