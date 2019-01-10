package domainHotelCommon

type CancelPenalty struct {
	HoursBefore int               `json:"hoursBefore"`
	Type        CancelPenaltyType `json:"penaltyType" gqlgen:"penaltyType"`
	Currency    string            `json:"currency"`
	// Zechao: ha cambiado de amount a value, borra el comentario
	Value float64 `json:"value"`
}
