package domainHotelCommon

type PriceBreakDown struct {
	EffectiveDate string `json:"effectiveDate"`
	ExpireDate    string `json:"expireDate"`
	Price         Price  `json:"price"`
}
