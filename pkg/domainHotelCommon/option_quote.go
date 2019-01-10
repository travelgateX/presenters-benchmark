package domainHotelCommon

type OptionQuote struct {
	OptionRefId  string
	Status       StatusType
	Price        Price
	CancelPolicy CancelPolicy
	Remarks      *string
	Surcharges   []Surcharge
	PaymentCards *[]string
	OptionId     *OptionID
	Supplier     string
}

func (o *OptionQuote) Prices() []*Price {
	prices := make([]*Price, 0, 1)
	// Price
	prices = append(prices, &o.Price)

	return prices
}
