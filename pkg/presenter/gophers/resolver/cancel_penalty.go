package graphResolver

import (
	"fmt"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
	"strconv"
)

type CancelPenaltyResolver struct {
	CancelPenalty *domainHotelCommon.CancelPenalty
}

func (r *CancelPenaltyResolver) HoursBefore() int32 {
	return int32(r.CancelPenalty.HoursBefore)
}

func (r *CancelPenaltyResolver) PenaltyType() domainHotelCommon.CancelPenaltyType {
	return (*r.CancelPenalty).Type
}

func (r *CancelPenaltyResolver) Currency() Currency {
	return Currency((*r.CancelPenalty).Currency)
}

func (r *CancelPenaltyResolver) Value() float64 {
	return r.CancelPenalty.Value
}

type Currency string

func (_ Currency) ImplementsGraphQLType(name string) bool {
	return name == "Currency"
}

func (currency *Currency) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*currency = Currency(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (currency Currency) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(currency)), nil
}
