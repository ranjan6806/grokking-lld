package repository

type URLRepository interface {
	Save(shortURL, originalURL string) error
	Find(shortURL string) (string, error)
}
