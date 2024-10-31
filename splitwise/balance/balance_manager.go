package balance

import (
	"splitwise/models"
	"splitwise/repository"
)

type BalanceManager struct {
	BalanceRepo repository.BalanceRepositoryInterface
}

func (bm *BalanceManager) AddUserBalance(user *models.User) {
	bm.BalanceRepo.SaveBalance(user.ID, make(map[string]float64))
}

func (bm *BalanceManager) AddExpense(expense map[string]float64, payer *models.User) {
	for userID, amount := range expense {
		if userID == payer.ID {
			continue
		}

		payerBalance := bm.BalanceRepo.GetBalance(payer.ID)
		receiverBalance := bm.BalanceRepo.GetBalance(userID)

		payerBalance[userID] += amount
		receiverBalance[payer.ID] -= amount

		bm.BalanceRepo.UpdateBalance(payer.ID, payerBalance)
		bm.BalanceRepo.UpdateBalance(userID, receiverBalance)
	}
}

func (bm *BalanceManager) ShowBalances() map[string]map[string]float64 {
	balances := make(map[string]map[string]float64)
	for _, userID := range bm.BalanceRepo.GetAllUsers() {
		balance := bm.BalanceRepo.GetBalance(userID)
		balances[userID] = balance
	}
	return balances
}

func (bm *BalanceManager) SettleBetweenUsers(user1ID, user2ID string) {
	// Retrieve current balances for both users
	user1Balance := bm.BalanceRepo.GetBalance(user1ID)
	user2Balance := bm.BalanceRepo.GetBalance(user2ID)

	user1Balance[user2ID] = 0
	user2Balance[user1ID] = 0

	bm.BalanceRepo.UpdateBalance(user1ID, user1Balance)
	bm.BalanceRepo.UpdateBalance(user2ID, user2Balance)
}

func (bm *BalanceManager) AddGroupExpense(groupID string, expense map[string]float64, payer *models.User, members []*models.User) {
	totalMembers := len(members)
	splitAmount := expense[payer.ID] / float64(totalMembers)

	for _, member := range members {
		memberID := member.ID
		if memberID == payer.ID {
			continue
		}

		// Retrieve current group balance
		payerGroupBalance := bm.BalanceRepo.GetGroupBalance(groupID, payer.ID)
		memberGroupBalance := bm.BalanceRepo.GetGroupBalance(groupID, memberID)

		// Update group balances for each user in the group
		payerGroupBalance[memberID] += splitAmount
		memberGroupBalance[payer.ID] -= splitAmount

		bm.BalanceRepo.UpdateGroupBalance(groupID, payer.ID, payerGroupBalance)
		bm.BalanceRepo.UpdateGroupBalance(groupID, memberID, memberGroupBalance)
	}
}

func (bm *BalanceManager) SettleGroup(groupID string) {
	groupUsers := bm.BalanceRepo.GetAllGroupUsers(groupID)

	for _, userID := range groupUsers {
		userBalance := bm.BalanceRepo.GetBalance(userID)

		for otherUserID := range userBalance {
			userBalance[otherUserID] = 0
		}

		// Update the cleared balance back in the repository
		bm.BalanceRepo.UpdateGroupBalance(groupID, userID, userBalance)
	}
}

func NewBalanceManager(balanceRepo repository.BalanceRepositoryInterface) BalanceManagerInterface {
	return &BalanceManager{
		BalanceRepo: balanceRepo,
	}
}
