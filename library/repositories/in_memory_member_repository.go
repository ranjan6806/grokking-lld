package repositories

import (
	"library/models"
	"sync"
)

type InMemoryMemberRepository struct {
	members map[string]*models.Member
	mu      sync.RWMutex
}

func (mr *InMemoryMemberRepository) Save(member *models.Member) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	mr.members[member.ID] = member
}

func (mr *InMemoryMemberRepository) RemoveMember(memberID string) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	delete(mr.members, memberID)
}

func (mr *InMemoryMemberRepository) FindByID(id string) (*models.Member, bool) {
	mr.mu.RLock()
	defer mr.mu.RUnlock()
	member, exists := mr.members[id]
	return member, exists
}

func (mr *InMemoryMemberRepository) IssueBook(memberID string, bookCopy *models.BookCopy) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	member, exists := mr.members[memberID]
	if !exists {
		return
	}

	if !member.CanIssueBook() {
		return
	}

	member.IssueBook(bookCopy)
}

func (mr *InMemoryMemberRepository) ReturnBook(memberID string, bookCopyID string) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	member, exists := mr.members[memberID]
	if !exists {
		return
	}

	member.ReturnBook(bookCopyID)
}

func NewInMemoryMemberRepository() MemberRepositoryInterface {
	return &InMemoryMemberRepository{
		members: make(map[string]*models.Member),
	}
}
