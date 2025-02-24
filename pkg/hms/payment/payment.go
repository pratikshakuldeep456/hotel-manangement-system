package payment

type Payment interface {
	Pay(amount int)
}

// methods, such as cash, credit card, and online payment.
