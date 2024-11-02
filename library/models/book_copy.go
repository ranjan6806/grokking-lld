package models

import "time"

type CopyStatus string

const (
	Available CopyStatus = "Available"
	Issued    CopyStatus = "Issued"
)

type BookCopy struct {
	ID         string
	BookID     string
	Status     CopyStatus
	IssuedTo   UserInterface
	IssuedDate *time.Time
}

func (bc *BookCopy) IsAvailable() bool {
	return bc.Status == Available
}

func (bc *BookCopy) Issue(member UserInterface) bool {
	if bc.IsAvailable() {
		bc.Status = Issued
		bc.IssuedTo = member
		now := time.Now()
		bc.IssuedDate = &now
		return true
	}

	return false
}

func (bc *BookCopy) Return() {
	bc.Status = Available
	bc.IssuedTo = nil
	bc.IssuedDate = nil
}
