package domainHotelCommon

type Supplement struct {
	Code           *string //Pointer because we need change value by the mapper
	Name           *string
	Description    *string
	SupplementType SupplementType
	ChargeType     ChargeType
	Mandatory      bool
	DurationType   *DurationType
	Quantity       *int
	Unit           *UnitTimeType
	EffectiveDate  *string
	ExpireDate     *string
	Resort         *Resort
	Price          *Price
}
