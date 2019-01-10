package domainHotelCommon

type CancelPolicy struct {
	Refundable      bool            `json:"refundable"`
	CancelPenalties []CancelPenalty `json:"cancelPenalties"`
}
