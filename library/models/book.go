package models

type Book struct {
	ID      string
	Title   string
	Author  string
	Subject string
	Copies  []*BookCopy
}

func (b *Book) AvailableCopies() int {
	count := 0
	for _, bookCopy := range b.Copies {
		if bookCopy.IsAvailable() {
			count++
		}
	}
	return count
}
