package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
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

func (r *PromotionResolver) EffectiveDate() *Date {
	if r.Promotion.EffectiveDate == nil {
		return nil
	}
	tmp := Date(*r.Promotion.EffectiveDate)
	return &tmp
}

func (r *PromotionResolver) ExpireDate() *Date {
	if r.Promotion.ExpireDate == nil {
		return nil
	}
	tmp := Date(*r.Promotion.ExpireDate)
	return &tmp
}
