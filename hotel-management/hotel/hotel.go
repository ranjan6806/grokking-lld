package hotel

import (
	"errors"
	"hotel-management/hotel_branch"
)

type Hotel struct {
	HotelID  string
	Name     string
	Branches map[string]*hotel_branch.HotelBranch // map string -> hotel branch
}

func NewHotel(id, name string) *Hotel {
	return &Hotel{
		HotelID:  id,
		Name:     name,
		Branches: make(map[string]*hotel_branch.HotelBranch),
	}
}

func (h *Hotel) AddBranch(b *hotel_branch.HotelBranch) error {
	_, exists := h.Branches[b.BranchID]
	if exists {
		return errors.New("hotel already exists")
	}

	h.Branches[b.BranchID] = b
	return nil
}

func (h *Hotel) GetBranch(branchID string) (*hotel_branch.HotelBranch, error) {
	if b, exists := h.Branches[branchID]; exists {
		return b, nil
	}

	return nil, errors.New("branch not found")
}

func (h *Hotel) RemoveBranch(branchID string) error {
	for _, branch := range h.Branches {
		if branch.BranchID == branchID {
			delete(h.Branches, branchID)
			return nil
		}
	}

	return errors.New("branch not found")
}
