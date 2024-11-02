package models

const MaxBorrowLimit = 10

type Member struct {
	ID          string
	Name        string
	LibraryCard string
	IssuedBooks []*BookCopy
}

func (m *Member) GetID() string {
	return m.ID
}

func (m *Member) GetName() string {
	return m.Name
}

func (m *Member) GetRole() Role {
	return RoleMember
}

func (m *Member) CanIssueBook() bool {
	return len(m.IssuedBooks) < MaxBorrowLimit
}

func (m *Member) IssueBook(bookCopy *BookCopy) {
	m.IssuedBooks = append(m.IssuedBooks, bookCopy)
}

func (m *Member) ReturnBook(bookCopyID string) {
	for i, bookCopy := range m.IssuedBooks {
		if bookCopy.ID == bookCopyID {
			m.IssuedBooks = append(m.IssuedBooks[:i], m.IssuedBooks[i+1:]...)
			break
		}
	}
}
