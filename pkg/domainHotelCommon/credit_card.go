package domainHotelCommon

type PaymentCard struct {
	CardType string
	Holder   Holder
	Number   string
	CVC      string
	Expire   PaymentCardExpiration
}

type PaymentCardExpiration struct {
	Month int32
	Year  int32
}
