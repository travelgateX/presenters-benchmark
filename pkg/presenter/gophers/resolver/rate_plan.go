package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type RatePlanResolver struct {
	RatePlan *domainHotelCommon.RatePlan
}

func (r *RatePlanResolver) Code() string {
	return *(*r.RatePlan).Code
}

func (r *RatePlanResolver) Name() *string {
	return (*r.RatePlan).Name
}

func (r *RatePlanResolver) EffectiveDate() *Date {
	if r.RatePlan.EffectiveDate == nil {
		return nil
	}
	tmp := Date(*r.RatePlan.EffectiveDate)
	return &tmp
}

func (r *RatePlanResolver) ExpireDate() *Date {
	if r.RatePlan.ExpireDate == nil {
		return nil
	}
	tmp := Date(*r.RatePlan.ExpireDate)
	return &tmp
}
