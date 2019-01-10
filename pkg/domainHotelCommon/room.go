package domainHotelCommon

type Room struct {
	OccupancyRefID int     `json:"occupancyRefId"`
	Code           *string `json:"code"` //Pointer because we need change value by the mapper
	OriginalCode   string
	Description    *string     `json:"description"`
	Refundable     *bool       `json:"refundable"`
	Units          *int        `json:"units"`
	RoomPrice      RoomPrice   `json:"roomPrice"`
	Beds           []Bed       `json:"beds"`
	RatePlans      []RatePlan  `json:"ratePlans"`
	Promotions     []Promotion `json:"promotions"`
}
