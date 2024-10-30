package controller

import (
	"parking-lot/models"
	"parking-lot/service"
)

type DisplayControllerInterface interface {
	ShowFreeSpots() map[models.SpotType]int
}

type DisplayController struct {
	displayService service.DisplayBoardServiceInterface
}

func NewDisplayController(displayService service.DisplayBoardServiceInterface) *DisplayController {
	return &DisplayController{
		displayService: displayService,
	}
}

func (dc *DisplayController) ShowFreeSpots() map[models.SpotType]int {
	return dc.displayService.ShowFreeSpots()
}
