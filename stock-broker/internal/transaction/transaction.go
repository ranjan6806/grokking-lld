package transaction

import "fmt"

type TransactionType string

const (
	CheckTransactionType    TransactionType = "check"
	WirelessTransactionType TransactionType = "wireless"
)

type Transaction struct {
	Amount float64
	Type   TransactionType
}

type TransactionStrategy interface {
	ProcessTransaction(t Transaction)
}

type CheckTransaction struct {
}

func (c *CheckTransaction) ProcessTransaction(t Transaction) {
	fmt.Printf("processing check transaction of $%.2f\n", t.Amount)
}
