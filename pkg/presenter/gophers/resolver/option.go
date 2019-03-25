package graphResolver

import (
	"github.com/travelgateX/presenters-benchmark/pkg/access"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
)

type HotelOptionResolver struct {
	Option *domainHotelCommon.Option

	context    *string
	group      *string
	rqDeepLink *string
	criteria   string
}

func (r *HotelOptionResolver) SupplierCode() string {
	return (*r.Option).Supplier
}

func (r *HotelOptionResolver) AccessCode() string {
	return (*r.Option).Access
}

func (r *HotelOptionResolver) Market() string {
	return r.Option.Market
}

func (r *HotelOptionResolver) HotelCode() string {
	return (*r.Option).HotelCode
}

func (r *HotelOptionResolver) HotelCodeSupplier() string {
	return (*r.Option).Id.HotelCode
}

func (r *HotelOptionResolver) HotelName() *string {
	if (*r.Option).HotelName != nil && *(*r.Option).HotelName == "" {
		return nil
	}
	return (*r.Option).HotelName
}

func (r *HotelOptionResolver) BoardCode() string {
	return *(*r.Option).BoardCode
}

func (r *HotelOptionResolver) BoardCodeSupplier() string {
	return (*r.Option).Id.BoardCode
}

func (r *HotelOptionResolver) PaymentType() string {
	return (*r.Option).PaymentType.Description()
}

func (r *HotelOptionResolver) Status() domainHotelCommon.StatusType {
	return (*r.Option).Status
}

func (r *HotelOptionResolver) Occupancies() []*OccupancyResolver {
	occupancies := make([]*OccupancyResolver, 0, len(r.Option.Occupancies))
	for _, occupancy := range r.Option.Occupancies {
		occupancy_aux := occupancy
		occupancies = append(occupancies, &OccupancyResolver{Occupancy: &occupancy_aux})
	}
	return occupancies
}

func (r *HotelOptionResolver) Rooms() []*RoomResolver {
	rooms := make([]*RoomResolver, 0, len(r.Option.Rooms))
	for _, room := range r.Option.Rooms {
		room_aux := room
		rooms = append(rooms, &RoomResolver{Room: &room_aux})
	}
	return rooms
}

func (r *HotelOptionResolver) Price() *PriceResolver {
	priceResolver := PriceResolver{Price: &(*r.Option).Price}
	return &priceResolver
}

func (r *HotelOptionResolver) Supplements() *[]*SupplementResolver {
	if len(r.Option.Supplements) < 1 {
		return nil
	}
	supplements := make([]*SupplementResolver, 0, len(r.Option.Supplements))
	for _, supplement := range r.Option.Supplements {
		supplement_aux := supplement
		supplements = append(supplements, &SupplementResolver{Supplement: supplement_aux})
	}
	return &supplements
}

func (r *HotelOptionResolver) Surcharges() *[]*SurchargeResolver {
	if r.Option.Surcharges == nil || len(r.Option.Surcharges) == 0 {
		return nil
	}
	surcharges := make([]*SurchargeResolver, 0, len(r.Option.Surcharges))
	for _, surcharge := range r.Option.Surcharges {
		surcharge_aux := surcharge
		surcharges = append(surcharges, &SurchargeResolver{Surcharge: &surcharge_aux})
	}
	return &surcharges
}

func (r *HotelOptionResolver) RateRules() *[]access.RateRulesType {
	if r.Option.RateRules == nil || len(r.Option.RateRules) == 0 {
		return nil
	}
	return &r.Option.RateRules
}

func (r *HotelOptionResolver) CancelPolicy() *CancelPolicyResolver {
	if (*r.Option).CancelPolicy == nil {
		return nil
	}
	CancelPolicyResolver := CancelPolicyResolver{CancelPolicy: &(*(*r.Option).CancelPolicy)}
	return &CancelPolicyResolver
}

func (r *HotelOptionResolver) Remarks() *string {
	return (*r.Option).Remarks
}

func (r *HotelOptionResolver) Token() string {
	return r.Option.Token
}

func (r *HotelOptionResolver) Id() string {
	return r.Option.OptionID
}

func (r *HotelOptionResolver) AddOns() *AddOnsResolver {
	return nil
}
