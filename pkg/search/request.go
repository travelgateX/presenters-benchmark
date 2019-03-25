package search

import (
	"github.com/travelgateX/presenters-benchmark/pkg/access"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type HotelSearchRQ struct {
	domainHotelCommon.HotelBaseRQ

	Token        *string         `gqlgen:"token"`
	Criteria     *CriteriaSearch `gqlgen:"criteria"`
	Filter       *Filter         `gqlgen:"filter"`
	FilterSearch *FilterSearch   `gqlgen:"filterSearch"`

	SerializedCriteria string
}

type CriteriaSearch struct {
	CheckIn      string
	CheckOut     string
	Hotels       *[]string
	Destinations *[]string
	Occupancies  []domainHotelCommon.Occupancy `gqlgen:"criteria"`
	Language     *string
	Currency     *string
	Nationality  *string
	Market       *string
}

type Filter struct {
	Access    *TypeFilter
	RateRules *RateRuleFilter
}

type FilterSearch struct {
	Access    *TypeFilter
	RateRules *RateRuleFilter
	Plugin    *domainHotelCommon.FilterPluginType
}

type TypeFilter struct {
	Includes *[]string
	Excludes *[]string
}

type RateRuleFilter struct {
	Includes *[]access.RateRulesType
	Excludes *[]access.RateRulesType
}
