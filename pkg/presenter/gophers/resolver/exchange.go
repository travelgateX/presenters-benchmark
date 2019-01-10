package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/domainHotelCommon"
)

type ExchangeResolver struct {
	Exchange domainHotelCommon.Exchange
}

func (r *ExchangeResolver) Currency() graphql.Currency {
	return graphql.Currency(r.Exchange.Currency)
}

func (r *ExchangeResolver) Rate() float64 {
	return r.Exchange.Rate
}
