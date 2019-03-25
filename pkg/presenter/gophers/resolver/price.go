package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type PriceResolver struct {
	Price *domainHotelCommon.Price
}

func (r *PriceResolver) Currency() Currency {
	return Currency(r.Price.Currency)
}

func (r *PriceResolver) Binding() bool {
	return r.Price.Binding
}

func (r *PriceResolver) Net() float64 {
	return r.Price.Net
}

func (r *PriceResolver) Gross() *float64 {
	return &r.Price.Gross
}

func (r *PriceResolver) Exchange() *ExchangeResolver {
	exchangeResolver := ExchangeResolver{Exchange: r.Price.Exchange}
	return &exchangeResolver
}

func (r *PriceResolver) Markups() *[]*MarkupResolver {
	if len(r.Price.Markups) == 0 {
		return nil
	}
	markups := make([]*MarkupResolver, 0, len(r.Price.Markups))
	for _, markup := range r.Price.Markups {
		markups = append(markups, &MarkupResolver{Markup: markup})
	}
	return &markups
}
