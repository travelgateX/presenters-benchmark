package graphResolver

import "github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"

type RoomResolver struct {
	Room *domainHotelCommon.Room
}

func (r *RoomResolver) OccupancyRefId() int32 {
	return int32(r.Room.OccupancyRefID)
}

func (r *RoomResolver) Code() string {
	return *(*r.Room).Code
}

func (r *RoomResolver) Description() *string {
	if (*r.Room).Description != nil && *(*r.Room).Description == "" {
		return nil
	}
	return (*r.Room).Description
}

func (r *RoomResolver) Refundable() *bool {
	return (*r.Room).Refundable
}

func (r *RoomResolver) Units() *int32 {
	if r.Room.Units == nil {
		return nil
	}
	u := int32(*r.Room.Units)
	return &u
}

func (r *RoomResolver) RoomPrice() *RoomPriceResolver {
	roomPriceResolver := RoomPriceResolver{RoomPrice: &(*r.Room).RoomPrice}
	return &roomPriceResolver
}

func (r *RoomResolver) Beds() *[]*BedResolver {
	if r.Room == nil || len(r.Room.Beds) == 0 {
		return nil
	}
	beds := make([]*BedResolver, 0, len(r.Room.Beds))
	for _, bed := range r.Room.Beds {
		bed_aux := bed
		beds = append(beds, &BedResolver{Bed: &bed_aux})
	}
	return &beds
}

func (r *RoomResolver) RatePlans() *[]*RatePlanResolver {
	if r.Room == nil || len(r.Room.RatePlans) == 0 {
		return nil
	}
	rate_plans := make([]*RatePlanResolver, 0, len(r.Room.RatePlans))
	for _, rate_plan := range r.Room.RatePlans {
		rate_plan_aux := rate_plan
		rate_plans = append(rate_plans, &RatePlanResolver{RatePlan: &rate_plan_aux})
	}
	return &rate_plans
}

func (r *RoomResolver) Promotions() *[]*PromotionResolver {
	if r.Room == nil || len(r.Room.Promotions) == 0 {
		return nil
	}
	promotions := make([]*PromotionResolver, 0, len(r.Room.Promotions))
	for _, promotion := range r.Room.Promotions {
		promotion_aux := promotion
		promotions = append(promotions, &PromotionResolver{Promotion: &promotion_aux})
	}
	return &promotions
}
