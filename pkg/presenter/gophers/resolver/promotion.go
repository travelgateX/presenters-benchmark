package graphResolver

import (
	"hub-aggregator/common/graphql"
	"rfc/presenters/pkg/domainHotelCommon"
)

type PromotionResolver struct {
	Promotion *domainHotelCommon.Promotion
}

func (r *PromotionResolver) Code() string {
	return (*r.Promotion).Code
}

func (r *PromotionResolver) Name() *string {
	return (*r.Promotion).Name
}

func (r *PromotionResolver) EffectiveDate() *graphql.Date {
	if r.Promotion.EffectiveDate == nil {
		return nil
	}
	tmp := graphql.Date(*r.Promotion.EffectiveDate)
	return &tmp
}

func (r *PromotionResolver) ExpireDate() *graphql.Date {
	if r.Promotion.ExpireDate == nil {
		return nil
	}
	tmp := graphql.Date(*r.Promotion.ExpireDate)
	return &tmp
}
