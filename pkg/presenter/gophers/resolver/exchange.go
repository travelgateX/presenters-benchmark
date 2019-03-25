package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type ExchangeResolver struct {
	Exchange domainHotelCommon.Exchange
}

func (r *ExchangeResolver) Currency() Currency {
	return Currency(r.Exchange.Currency)
}

func (r *ExchangeResolver) Rate() float64 {
	return r.Exchange.Rate
}
