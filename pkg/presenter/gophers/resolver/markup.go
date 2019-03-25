package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type MarkupResolver struct {
	Markup domainHotelCommon.Markup
}

func (r *MarkupResolver) Channel() *string {
	return r.Markup.Channel
}

func (r *MarkupResolver) Currency() Currency {
	return Currency(r.Markup.Currency)
}

func (r *MarkupResolver) Binding() bool {
	return r.Markup.Binding
}

func (r *MarkupResolver) Net() float64 {
	return r.Markup.Net
}

func (r *MarkupResolver) Gross() *float64 {
	return &r.Markup.Gross
}

func (r *MarkupResolver) Exchange() *ExchangeResolver {
	exchangeResolver := ExchangeResolver{Exchange: r.Markup.Exchange}
	return &exchangeResolver
}

func (r *MarkupResolver) Rules() []*RuleResolver {
	resolvers := make([]*RuleResolver, 0, len(r.Markup.Rules))
	for _, rule := range r.Markup.Rules {
		resolvers = append(resolvers, &RuleResolver{&rule})
	}
	return resolvers
}
