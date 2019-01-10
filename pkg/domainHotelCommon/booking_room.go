package domainHotelCommon

type BookingRoom struct {
	OccupancyRefId *int32
	Code           *string
	Description    *string
	Price          *Price
}
