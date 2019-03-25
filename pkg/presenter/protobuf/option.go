package protobuf

import (
	"github.com/travelgateX/presenters-benchmark/pkg/access"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
)

func NewResponse(options []*presenter.Option) *SearchReply {
	return &SearchReply{
		Options: newOptions(options),
	}
}

func newOptions(options []*presenter.Option) []*Option {
	ret := make([]*Option, 0, len(options))
	for key := range options {
		ret = append(ret, newOption(options[key]))
	}
	return ret
}

func newOption(option *presenter.Option) *Option {
	ret := Option{}
	ret.Id = option.OptionID
	ret.SupplierCode = option.Supplier
	ret.AccessCode = option.Access
	ret.Market = option.Market
	ret.HotelCode = option.HotelCode
	if option.HotelName != nil {
		ret.HotelName = *option.HotelName
	}
	if option.BoardCode != nil {
		ret.BoardCode = *option.BoardCode
	}
	ret.BoardCodeSupplier = option.BoardCodeOriginal
	ret.PaymentType = PaymentType(PaymentType_value[string(option.PaymentType)])
	ret.StatusType = StatusType_RQ
	if option.Remarks != nil {
		ret.Remarks = *option.Remarks
	}
	ret.Token = option.Token
	if len(option.Occupancies) != 0 {
		ret.Occupancies = make([]*Occupancy, 0, len(option.Occupancies))
		for key := range option.Occupancies {
			ret.Occupancies = append(ret.Occupancies, newOccupancy(option.Occupancies[key]))
		}
	}

	if len(option.Rooms) != 0 {
		ret.Rooms = make([]*Room, 0, len(option.Rooms))
		for key := range option.Rooms {
			ret.Rooms = append(ret.Rooms, newRoom(option.Rooms[key]))
		}
	}

	ret.Price = newPrice(&option.Price)
	if len(option.Supplements) != 0 {
		ret.Supplements = make([]*Supplement, 0, len(option.Supplements))
		for key := range option.Supplements {
			ret.Supplements = append(ret.Supplements, newSupplement(option.Supplements[key]))
		}
	}

	if len(option.Surcharges) != 0 {
		ret.Surcharges = make([]*Surcharge, 0, len(option.Surcharges))
		for key := range option.Surcharges {
			ret.Surcharges = append(ret.Surcharges, newSurcharges(option.Surcharges[key]))
		}
	}

	if len(option.RateRules) != 0 {
		ret.RateRules = make([]RateRuleType, 0, len(option.RateRules))
		for key := range option.RateRules {
			ret.RateRules = append(ret.RateRules, newRateRule(option.RateRules[key]))
		}
	}

	ret.CancelPolicy = newCancelPolicy(option.CancelPolicy)
	return &ret
}

func newCancelPolicy(policy *domainHotelCommon.CancelPolicy) *CancelPolicy {
	if policy == nil {
		return nil
	}
	ret := CancelPolicy{
		Refundable: policy.Refundable,
	}
	if policy != nil {
		if len(policy.CancelPenalties) != 0 {
			ret.CancelPenalties = make([]*CancelPenalty, 0, len(policy.CancelPenalties))
			for key := range policy.CancelPenalties {
				ret.CancelPenalties = append(ret.CancelPenalties, newCancelPenalty(policy.CancelPenalties[key]))
			}
		}
	}
	return &ret
}

func newCancelPenalty(penalty domainHotelCommon.CancelPenalty) *CancelPenalty {
	ret := CancelPenalty{
		HoursBefore: int64(penalty.HoursBefore),
		PenaltyType: CancelPenaltyType(CancelPenaltyType_value[string(penalty.Type)]),
		Currency:    penalty.Currency,
		Value:       penalty.Value,
	}
	return &ret
}

func newRateRule(rulesType access.RateRulesType) RateRuleType {
	return RateRuleType(RateRuleType_value[string(rulesType)])
}

func newSurcharges(surcharge domainHotelCommon.Surcharge) *Surcharge {
	var desc string
	if surcharge.Description != nil {
		desc = *surcharge.Description
	}
	ret := Surcharge{
		ChargeType:  ChargeType(ChargeType_value[string(surcharge.ChargeType)]),
		Mandatory:   surcharge.Mandatory,
		Description: desc,
		Price:       newPrice(&surcharge.Price),
	}
	return &ret
}

func newSupplement(supplement *domainHotelCommon.Supplement) *Supplement {
	if supplement == nil {
		return nil
	}
	var name, desc string
	if supplement.Name != nil {
		name = *supplement.Name
	}
	if supplement.Description != nil {
		desc = *supplement.Description
	}
	ret := Supplement{
		Name:           name,
		Description:    desc,
		SupplementType: SupplementType(SupplementType_value[string(supplement.SupplementType)]),
		ChargeType:     ChargeType(ChargeType_value[string(supplement.ChargeType)]),
		Mandatory:      supplement.Mandatory,
	}
	if supplement.Code == nil {
		ret.Code = ""
	} else {
		ret.Code = *supplement.Code
	}
	if supplement.DurationType != nil {
		ret.DurationType = DurationType(DurationType_value[string(*supplement.DurationType)])
	}
	if supplement.Quantity != nil {
		ret.Quantity = int64(*supplement.Quantity)
	}
	if supplement.Unit != nil {
		ret.UnitTimeType = UnitTimeType(UnitTimeType_value[string(*supplement.Unit)])
	}
	if supplement.EffectiveDate != nil {
		ret.EffectiveDate = *supplement.EffectiveDate
	}
	if supplement.ExpireDate != nil {
		ret.ExpireDate = *supplement.ExpireDate
	}
	ret.Resort = newResort(supplement.Resort)
	ret.Price = newPrice(supplement.Price)
	return &ret
}

func newResort(resort *domainHotelCommon.Resort) *Resort {
	if resort == nil {
		return nil
	}

	var name, desc string
	if resort.Name != nil {
		name = *resort.Name
	}
	if resort.Description != nil {
		desc = *resort.Description
	}
	ret := Resort{
		Code:        resort.Code,
		Name:        name,
		Description: desc,
	}
	return &ret
}

func newOccupancy(occupancy domainHotelCommon.Occupancy) *Occupancy {
	ret := Occupancy{
		Id: int32(occupancy.Id),
	}

	if len(occupancy.Paxes) != 0 {
		ret.Paxes = make([]*Pax, 0, len(occupancy.Paxes))
		for key := range occupancy.Paxes {
			ret.Paxes = append(ret.Paxes, newPax(occupancy.Paxes[key]))
		}
	}
	return &ret
}

func newPax(pax domainHotelCommon.Pax) *Pax {
	return &Pax{
		Age: uint32(pax.Age),
	}
}

func newRoom(room domainHotelCommon.Room) *Room {
	var desc string
	var ref bool
	var units int64
	if room.Description != nil {
		desc = *room.Description
	}
	if room.Refundable != nil {
		ref = *room.Refundable
	}
	if room.Units != nil {
		units = int64(*room.Units)
	}
	ret := &Room{
		OccupancyRefId: int32(room.OccupancyRefID),
		Description:    desc,
		Refundable:     ref,
		Units:          units,
	}
	if room.Code == nil {
		ret.Code = ""
	} else {
		ret.Code = *room.Code
	}
	ret.RoomPrice = newRoomPrice(room.RoomPrice)
	if len(room.Beds) != 0 {
		ret.Beds = make([]*Bed, 0, len(room.Beds))
		for key := range room.Beds {
			ret.Beds = append(ret.Beds, newBed(room.Beds[key]))
		}
	}
	if len(room.RatePlans) != 0 {
		ret.RatePlans = make([]*RatePlan, 0, len(room.RatePlans))
		for key := range room.RatePlans {
			ret.RatePlans = append(ret.RatePlans, newRatePlan(room.RatePlans[key]))
		}
	}
	if len(room.Promotions) != 0 {
		ret.Promotions = make([]*Promotion, 0, len(room.Promotions))
		for key := range room.Promotions {
			ret.Promotions = append(ret.Promotions, newPromotion(room.Promotions[key]))
		}
	}
	return ret
}

func newPromotion(promotion domainHotelCommon.Promotion) *Promotion {
	ret := &Promotion{
		Code: promotion.Code,
	}
	if promotion.Name != nil {
		ret.Name = *promotion.Name
	}
	if promotion.EffectiveDate != nil {
		ret.EffectiveDate = *promotion.EffectiveDate
	}
	if promotion.ExpireDate != nil {
		ret.ExpireDate = *promotion.ExpireDate
	}
	return ret
}

func newRatePlan(plan domainHotelCommon.RatePlan) *RatePlan {
	ret := RatePlan{}
	if plan.Name != nil {
		ret.Name = *plan.Name
	}
	if plan.EffectiveDate != nil {
		ret.EffectiveDate = *plan.EffectiveDate
	}
	if plan.ExpireDate != nil {
		ret.ExpireDate = *plan.ExpireDate
	}
	if plan.Code != nil {
		ret.Code = *plan.Code
	}
	return &ret
}

func newBed(bed domainHotelCommon.Bed) *Bed {
	ret := Bed{}
	if bed.Type != nil {
		ret.Type = *bed.Type
	}
	if bed.Description != nil {
		ret.Description = *bed.Description
	}
	if bed.Count != nil {
		ret.Count = int64(*bed.Count)
	}
	if bed.Shared != nil {
		ret.Shared = *bed.Shared
	}
	return &ret
}

func newRoomPrice(price domainHotelCommon.RoomPrice) *RoomPrice {
	ret := RoomPrice{}
	ret.Price = newPrice(&price.Price)
	if len(price.Breakdown) != 0 {
		ret.Breakdowns = make([]*Breakdown, 0, len(price.Breakdown))
		for key := range price.Breakdown {
			ret.Breakdowns = append(ret.Breakdowns, newPriceBreakdown(price.Breakdown[key]))
		}
	}
	return &ret
}

func newPriceBreakdown(down domainHotelCommon.PriceBreakDown) *Breakdown {
	ret := Breakdown{
		EffectiveDate: down.EffectiveDate,
		ExpireDate:    down.ExpireDate,
		Price:         newPrice(&down.Price),
	}
	return &ret
}

func newPrice(price *domainHotelCommon.Price) *Price {
	if price == nil {
		return nil
	}
	ret := Price{
		Currency: price.Currency,
		Binding:  price.Binding,
		Net:      price.Net,
		Gross:    price.Gross,
		Exchange: newExchange(price.Exchange),
	}
	if len(price.Markups) != 0 {
		ret.Markups = make([]*Markup, 0, len(price.Markups))
		for key := range price.Markups {
			ret.Markups = append(ret.Markups, newMarkup(price.Markups[key]))
		}
	}
	return &ret
}

func newMarkup(markup domainHotelCommon.Markup) *Markup {
	ret := Markup{
		Currency: markup.Currency,
		Binding:  markup.Binding,
		Net:      markup.Net,
		Gross:    markup.Gross,
		Exchange: newExchange(markup.Exchange),
	}
	if markup.Channel != nil {
		ret.Channel = *markup.Channel
	}
	if len(markup.Rules) != 0 {
		ret.Rules = make([]*Rule, 0, len(markup.Rules))
		for key := range markup.Rules {
			ret.Rules = append(ret.Rules, newRule(markup.Rules[key]))
		}
	}
	return &ret
}

func newRule(rule domainHotelCommon.Rule) *Rule {
	ret := Rule{
		Id:    rule.Id,
		Type:  Rule_MarkupRuleType(Rule_MarkupRuleType_value[string(rule.Type)]),
		Value: rule.Value,
	}
	if rule.Name != nil {
		ret.Name = *rule.Name
	}
	return &ret
}

func newExchange(exchange domainHotelCommon.Exchange) *Exchange {
	ret := Exchange{
		Currency: exchange.Currency,
		Rate:     exchange.Rate,
	}
	return &ret
}
