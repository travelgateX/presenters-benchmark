package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type SupplementResolver struct {
	Supplement *domainHotelCommon.Supplement
}

func (r *SupplementResolver) Code() string {
	return *(*r.Supplement).Code
}

func (r *SupplementResolver) Name() *string {
	return (*r.Supplement).Name
}

func (r *SupplementResolver) Description() *string {
	return (*r.Supplement).Description
}

func (r *SupplementResolver) SupplementType() domainHotelCommon.SupplementType {
	return (*r.Supplement).SupplementType
}

func (r *SupplementResolver) ChargeType() domainHotelCommon.ChargeType {
	return (*r.Supplement).ChargeType
}

func (r *SupplementResolver) Mandatory() bool {
	return (*r.Supplement).Mandatory
}

func (r *SupplementResolver) DurationType() *domainHotelCommon.DurationType {
	return (*r.Supplement).DurationType
}

func (r *SupplementResolver) Quantity() *int32 {
	if r.Supplement.Quantity == nil {
		return nil
	}
	q := int32(*r.Supplement.Quantity)
	return &q
}

func (r *SupplementResolver) Unit() *domainHotelCommon.UnitTimeType {
	return (*r.Supplement).Unit
}

func (r *SupplementResolver) EffectiveDate() *Date {
	if (*r.Supplement).EffectiveDate == nil {
		return nil
	}
	tmp := Date(*(*r.Supplement).EffectiveDate)
	return &tmp
}

func (r *SupplementResolver) ExpireDate() *Date {
	if (*r.Supplement).ExpireDate == nil {
		return nil
	}
	tmp := Date(*(*r.Supplement).ExpireDate)
	return &tmp
}

func (r *SupplementResolver) Resort() *ResortResolver {
	if (*r.Supplement).Resort == nil {
		return nil
	}
	resortResolver := ResortResolver{Resort: &(*(*r.Supplement).Resort)}
	return &resortResolver
}

func (r *SupplementResolver) Price() *PriceResolver {
	if (*r.Supplement).Price == nil {
		return nil
	}
	priceResolver := PriceResolver{Price: &(*(*r.Supplement).Price)}
	return &priceResolver
}
