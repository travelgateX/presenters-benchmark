package domainHotelCommon

type Surcharge struct {
	ChargeType  ChargeType
	Mandatory   bool
	Price       Price
	Description *string
}
