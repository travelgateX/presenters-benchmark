package domainHotelCommon

type Occupancy struct {
	Id    int   `json:"id"`
	Paxes []Pax `json:"paxes"`
}
