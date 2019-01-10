package domainHotelCommon

type CancelDetail struct {
	Reference             *Reference
	CancellationReference *string
	Status                *BookingStatusType
	Price                 *Price
}

func (cd *CancelDetail) Prices() []*Price {
	if cd.Price == nil {
		return []*Price{}
	}
	prices := make([]*Price, 0, 1)
	prices = append(prices, cd.Price)
	return prices
}
