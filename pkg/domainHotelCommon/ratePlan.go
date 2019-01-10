package domainHotelCommon

type RatePlan struct {
	Code          *string //Pointer because we need change value by the mapper
	OriginalCode  string
	Name          *string
	EffectiveDate *string
	ExpireDate    *string
}
