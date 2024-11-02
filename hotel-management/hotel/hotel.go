package hotel

import (
	"errors"
	"hotel-management/hotel_branch"
)

type Hotel struct {
	HotelID  string
	Name     string
	Branches []*hotel_branch.HotelBranch
}

func NewHotel(id, name string) *Hotel {
	return &Hotel{
		HotelID:  id,
		Name:     name,
		Branches: make([]*hotel_branch.HotelBranch, 0),
	}
}

func (h *Hotel) AddBranch(b *hotel_branch.HotelBranch) error {
	h.Branches = append(h.Branches, b)
	return nil
}

func (h *Hotel) RemoveBranch(branchID string) error {
	for i, branch := range h.Branches {
		if branch.BranchID == branchID {
			h.Branches = append(h.Branches[:i], h.Branches[i+1:]...)
			return nil
		}
	}

	return errors.New("branch not found")
}
