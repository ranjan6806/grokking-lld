package repository

import (
	"fmt"
	"splitwise/models"
)

type InMemoryGroupRepository struct {
	Groups       map[string]*models.Group
	groupCounter int
}

func (gr *InMemoryGroupRepository) AddGroup(group *models.Group) {
	groupID := fmt.Sprintf("g%d", gr.groupCounter)
	group.ID = groupID
	gr.Groups[groupID] = group
	gr.groupCounter++
}

func (gr *InMemoryGroupRepository) RemoveGroup(groupID string) {
	delete(gr.Groups, groupID)
}

func (gr *InMemoryGroupRepository) GetGroup(groupID string) *models.Group {
	return gr.Groups[groupID]
}

func NewGroupRepository() GroupRepositoryInterface {
	return &InMemoryGroupRepository{
		Groups:       make(map[string]*models.Group),
		groupCounter: 1,
	}
}
