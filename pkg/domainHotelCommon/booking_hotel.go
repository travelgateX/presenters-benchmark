package domainHotelCommon

type BookingHotel struct {
	CreationDate *string
	CheckIn      *string
	CheckOut     *string
	HotelCode    *string
	HotelName    *string
	BoardCode    *string
	Occupancies  *[]Occupancy
	Rooms        *[]BookingRoom
	//CancelPolicy *CancelPolicy
	//Remarks      *string
	//Status       BookingStatusType
}
