package payment

import "fmt"

type Payment interface {
	Pay(amount int) bool
}

type CashPayment struct {
}

func NewCashPayment() *CashPayment {
	return &CashPayment{}
}

func (cp *CashPayment) Pay(amount int) bool {
	fmt.Println("payment is done")
	return true
}

type CardPayment struct {
}

func NewCardPayment() *CardPayment {
	return &CardPayment{}
}
func (cp *CardPayment) Pay(amount int) bool {
	fmt.Println("payment is done")
	return true
}

// methods, such as cash, credit card, and online payment.
