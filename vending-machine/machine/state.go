package machine

type State interface {
	InsertMoney(amount uint)
	SelectProduct(row, column uint)
	DispenseProduct(row, column uint)
}
