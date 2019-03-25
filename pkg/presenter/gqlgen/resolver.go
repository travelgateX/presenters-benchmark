package gqlgen

import (
	"context"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
)

type Resolver struct {
	Options []*presenter.Option
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) HotelX(ctx context.Context) (*HotelXQuery, error) {
	ret := make([]HotelOptionSearch, 0, len(r.Options))
	for i := range r.Options {
		o := r.Options[i]

		// Cancellation Policy
		var cp *CancelPolicy
		if o.CancelPolicy != nil {
			cp = new(CancelPolicy)
			cp.Refundable = o.CancelPolicy.Refundable
			cp.CancelPenalties = make([]CancelPenalty, 0, len(o.CancelPolicy.CancelPenalties))
			for _, cpen := range o.CancelPolicy.CancelPenalties {
				cp.CancelPenalties = append(cp.CancelPenalties, CancelPenalty{
					Currency:    cpen.Currency,
					HoursBefore: cpen.HoursBefore,
					PenaltyType: CancelPenaltyType(cpen.Type.String()),
					Value:       cpen.Value,
				})
			}
		}

		// Occupancies
		var occupancies []Occupancy
		if len(o.Occupancies) != 0 {
			occupancies = make([]Occupancy, 0, len(o.Occupancies))
			for _, oc := range o.Occupancies {
				paxes := make([]Pax, 0, len(oc.Paxes))
				for _, pax := range oc.Paxes {
					paxes = append(paxes, Pax{Age: pax.Age})
				}
				occupancies = append(occupancies, Occupancy{
					ID:    oc.Id,
					Paxes: paxes,
				})
			}
		}

		//MARKUPS
		var mm []Markup
		for mi := range o.Price.Markups {
			rr := make([]Rule, 0, len(o.Price.Markups[mi].Rules))
			for ri := range o.Price.Markups[i].Rules {
				rr = append(rr, toRule(o.Price.Markups[mi].Rules[ri]))
			}
			mm = append(mm, toMarkup(o.Price.Markups[i]))
		}

		//SURCHARGES
		var surcharges []Surcharge
		if len(o.Surcharges) != 0 {
			surcharges = make([]Surcharge, 0, len(o.Surcharges))
			for si := range o.Surcharges {
				su := o.Surcharges[si]
				surcharges = append(surcharges, Surcharge{
					Price:       toPrice(su.Price),
					Description: su.Description,
					ChargeType:  ChargeType(su.ChargeType),
					Mandatory:   su.Mandatory,
				})
			}
		}
		//SUPPLEMENTS
		var supplements []Supplement
		if len(o.Supplements) != 0 {
			supplements = make([]Supplement, 0, len(o.Supplements))
			for si := range o.Supplements {
				su := o.Supplements[si]

				p := toPrice(*su.Price)
				dt := DurationType(*su.DurationType)
				u := UnitTimeType(*su.Unit)
				q := int(*su.Quantity)
				supplements = append(supplements, Supplement{
					Mandatory:     su.Mandatory,
					ChargeType:    ChargeType(su.ChargeType),
					Description:   su.Description,
					Price:         &p,
					Name:          su.Name,
					Code:          *su.Code,
					DurationType:  &dt,
					EffectiveDate: su.EffectiveDate,
					ExpireDate:    su.EffectiveDate,
					Quantity:      &q,
					Resort: &Resort{
						Code:        su.Resort.Code,
						Name:        su.Resort.Name,
						Description: su.Resort.Description,
					},
					SupplementType: SupplementType(su.SupplementType),
					Unit:           &u,
				})
			}
		}

		rooms := make([]Room, 0, len(o.Rooms))
		for ri := range o.Rooms {
			rooms = append(rooms, toRoom(o.Rooms[ri]))
		}

		var rateRules []RateRulesType
		if o.RateRules != nil {
			rateRules = make([]RateRulesType, 0, len(o.RateRules))
			for rri := range o.RateRules {
				rr := o.RateRules[rri]
				rateRules = append(rateRules, RateRulesType(rr))
			}
		}
		newOpt := HotelOptionSearch{
			Price:        toPrice(o.Price),
			ID:           o.OptionID,
			Remarks:      o.Remarks,
			Occupancies:  occupancies,
			CancelPolicy: cp,
			//PaymentType:       PaymentType(o.PaymentType),
			SupplierCode:      o.Supplier,
			Market:            o.Market,
			Status:            StatusType(o.Status),
			AccessCode:        o.Access,
			BoardCode:         o.BoardCodeOriginal,
			BoardCodeSupplier: *o.BoardCode,
			HotelCode:         o.HotelCode,
			HotelName:         o.HotelName,
			Surcharges:        nil,
			Supplements:       nil,
			Rooms:             rooms,
			RateRules:         nil,
			Token:             o.Token,
		}
		ret = append(ret, newOpt)

	}
	return &HotelXQuery{
		Search: &HotelSearch{
			Options: ret,
		},
	}, nil
}
func (r *queryResolver) Search(ctx context.Context) (Search, error) {
	panic("not implemented")
}

func toMarkup(m domainHotelCommon.Markup) Markup {
	rules := make([]Rule, 0, len(m.Rules))
	for i := range m.Rules {
		rules = append(rules, toRule(m.Rules[i]))
	}

	return Markup{
		Currency: m.Currency,
		Binding:  m.Binding,
		Gross:    &m.Gross,
		Net:      m.Net,
		Channel:  m.Channel,
		Exchange: Exchange{
			Currency: m.Exchange.Currency,
			Rate:     m.Exchange.Rate,
		},
		Rules: rules,
	}
}
func toRule(rule domainHotelCommon.Rule) Rule {
	return Rule{
		Type:  MarkupRuleType(rule.Type),
		Value: rule.Value,
		Name:  rule.Name,
		ID:    rule.Id,
	}
}

func toPrice(price domainHotelCommon.Price) Price {
	mu := make([]Markup, 0, len(price.Markups))
	for i := range price.Markups {
		mu = append(mu, toMarkup(price.Markups[i]))
	}

	return Price{
		Markups: mu,
		Exchange: Exchange{
			Currency: price.Exchange.Currency,
			Rate:     price.Exchange.Rate,
		},
		Currency: price.Currency,
		Net:      price.Net,
		Gross:    &price.Gross,
		Binding:  price.Binding,
	}
}

func toRoom(room domainHotelCommon.Room) Room {

	var beds []Bed
	if room.Beds != nil {
		beds = make([]Bed, 0, len(room.Beds))
		for i := range room.Beds {
			beds = append(beds, toBed(room.Beds[i]))
		}
	}

	var promotions []Promotion
	if room.Promotions != nil {
		promotions = make([]Promotion, 0, len(room.Promotions))
		for i := range room.Promotions {
			promotions = append(promotions, toPromotion(room.Promotions[i]))
		}
	}

	var ratePlans []RatePlan
	if room.RatePlans != nil {
		ratePlans = make([]RatePlan, 0, len(room.RatePlans))
		for i := range room.RatePlans {
			ratePlans = append(ratePlans, toRatePlan(room.RatePlans[i]))
		}
	}

	var u int
	if room.Units != nil {
		u = int(*room.Units)
	}
	return Room{
		Description:    room.Description,
		Code:           *room.Code,
		RoomPrice:      toRoomPrice(room.RoomPrice),
		OccupancyRefID: int(room.OccupancyRefID),
		Refundable:     room.Refundable,
		Units:          &u,
		Beds:           beds,
		Promotions:     promotions,
		RatePlans:      ratePlans,
	}
}

func toRoomPrice(price domainHotelCommon.RoomPrice) RoomPrice {
	var bd []PriceBreakdown
	if price.Breakdown != nil {
		bd = make([]PriceBreakdown, 0, len(price.Breakdown))
		for i := range price.Breakdown {
			bd = append(bd, PriceBreakdown{
				Price:         toPrice(price.Breakdown[i].Price),
				EffectiveDate: price.Breakdown[i].EffectiveDate,
				ExpireDate:    price.Breakdown[i].ExpireDate,
			})
		}
	}
	return RoomPrice{
		Price:     toPrice(price.Price),
		Breakdown: bd,
	}
}

func toBed(bed domainHotelCommon.Bed) Bed {
	var c int
	if bed.Count != nil {
		c = int(*bed.Count)
	}
	return Bed{
		Type:        bed.Type,
		Description: bed.Description,
		Count:       &c,
		Shared:      bed.Shared,
	}
}

func toPromotion(promotion domainHotelCommon.Promotion) Promotion {
	return Promotion{
		Code:          promotion.Code,
		Name:          promotion.Name,
		ExpireDate:    promotion.ExpireDate,
		EffectiveDate: promotion.EffectiveDate,
	}
}

func toRatePlan(ratePlan domainHotelCommon.RatePlan) RatePlan {
	return RatePlan{
		EffectiveDate: ratePlan.EffectiveDate,
		ExpireDate:    ratePlan.ExpireDate,
		Name:          ratePlan.Name,
		Code:          *ratePlan.Code,
	}
}
