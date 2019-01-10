package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/domainHotelCommon"
)

type PriceResolver struct {
	Price *domainHotelCommon.Price
}

func (r *PriceResolver) Currency() graphql.Currency {
	return graphql.Currency(r.Price.Currency)
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
