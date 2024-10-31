package expenses

import "splitwise/models"

type SplitStrategyInterface interface {
	CalculateSplit(amount float64, users []*models.User) map[string]float64
}

type EqualSplit struct {
}

func (e *EqualSplit) CalculateSplit(amount float64, users []*models.User) map[string]float64 {
	split := amount / float64(len(users))
	result := make(map[string]float64)
	for _, user := range users {
		result[user.ID] = split
	}
	return result
}

type ExactSplit struct {
	Splits map[string]float64
}

func (e *ExactSplit) CalculateSplit(amount float64, users []*models.User) map[string]float64 {
	return e.Splits
}
