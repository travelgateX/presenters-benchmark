package domainHotelCommon

type RoomPrice struct {
	Price     Price            `json:"price"`
	Breakdown []PriceBreakDown `json:"breakdowns"`
}
