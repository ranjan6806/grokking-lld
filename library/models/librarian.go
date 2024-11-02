package models

type Librarian struct {
	ID         string
	Name       string
	EmployeeID string
}

func (l *Librarian) GetID() string {
	return l.ID
}

func (l *Librarian) GetName() string {
	return l.Name
}

func (l *Librarian) GetRole() Role {
	return RoleLibrarian
}
