package graphResolver

import (
	"context"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
	"github.com/travelgateX/presenters-benchmark/pkg/search"
)

type SearchResolver struct {
}

type HotelSearchRQ struct {
	Token        *string
	Criteria     *search.CriteriaSearch
	Filter       *search.Filter
	FilterSearch *search.FilterSearch
	Settings     *struct {
		Context           *string
		Client            *string
		Group             *string
		Org               string
		Suppliers         *[]domainHotelCommon.Supplier
		Plugins           *[]domainHotelCommon.PluginStep
		TestMode          *bool
		ClientTokens      *[]string
		Timeout           *int
		AuditTransactions *bool
		BusinessRules     *struct {
			OptionsQuota      *int
			BusinessRulesType *string
		}
		Currency *string

		UseContext  *string
		ConnectUser *string
	}
}

func (r *SearchResolver) Hotel(ctx context.Context) *HotelSearchResolver {
	return nil
}
