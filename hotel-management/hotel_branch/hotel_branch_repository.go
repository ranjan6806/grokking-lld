package hotel_branch

import (
	"fmt"
	"sync"
)

type HotelBranchRepository interface {
	AddBranch(branch *HotelBranch) error
	GetBranch(id string) (*HotelBranch, error)
}

type HotelBranchRepositoryImpl struct {
	branches map[string]*HotelBranch
	mtx      sync.RWMutex
}

func NewHotelBranchRepository() HotelBranchRepository {
	return &HotelBranchRepositoryImpl{
		branches: make(map[string]*HotelBranch),
	}
}

func (r *HotelBranchRepositoryImpl) AddBranch(branch *HotelBranch) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.branches[branch.BranchID] = branch
	return nil
}

func (r *HotelBranchRepositoryImpl) GetBranch(id string) (*HotelBranch, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	branch, ok := r.branches[id]
	if !ok {
		return nil, fmt.Errorf("hotel branch with id %s not found", id)
	}

	return branch, nil
}
