package requests

type OrderCreate struct {
	UserID uint
}

type OrderUpdate struct {
	ID              uint
	RefCode         string
	Amount          float64
	SlipImage       string
	UserID          uint
	PaymentMethodID uint
}
