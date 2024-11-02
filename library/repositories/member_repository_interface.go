package repositories

import "library/models"

type MemberRepositoryInterface interface {
	Save(member *models.Member)
	RemoveMember(memberID string)
	FindByID(id string) (*models.Member, bool)
	IssueBook(memberID string, bookCopy *models.BookCopy)
	ReturnBook(memberID string, bookCopyID string)
}
