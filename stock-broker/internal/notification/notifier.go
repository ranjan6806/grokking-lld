package notification

import (
	"fmt"
	"stock-broker/internal/order"
)

type Notifier interface {
	Notify(order order.Order)
}

type EmailNotifier struct{}

func (en *EmailNotifier) Notify(order order.Order) {
	fmt.Println("Sending email notification for order: ", order.Execute())
}
