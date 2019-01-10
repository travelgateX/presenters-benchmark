package domainHotelCommon

type BookingDetail struct {
	Reference    Reference
	Holder       *Holder
	Hotel        *BookingHotel
	Price        *Price
	CancelPolicy *CancelPolicy
	Remarks      *string
	Status       BookingStatusType
	Payable      *string
	Supplier     string
	OptionIDRq   *OptionID
}

func (o *BookingDetail) Prices() []*Price {
	prices := make([]*Price, 0, 5)

	// Price
	if o.Price != nil {
		prices = append(prices, o.Price)
	}

	// Rooms
	if o.Hotel != nil && o.Hotel.Rooms != nil && len(*o.Hotel.Rooms) > 0 {
		for iRoom := range *o.Hotel.Rooms {
			if (*o.Hotel.Rooms)[iRoom].Price != nil {
				prices = append(prices, (*o.Hotel.Rooms)[iRoom].Price)
			}
		}
	}

	return prices
}
