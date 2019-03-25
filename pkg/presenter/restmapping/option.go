package restmapping

import (
	"github.com/travelgateX/presenters-benchmark/pkg/access"
	"github.com/travelgateX/presenters-benchmark/pkg/domainHotelCommon"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
)

type Response struct {
	Data struct {
		HotelX struct {
			Search struct {
				Options []Option `json:"options"`
				Errors  struct {
					Code        string `json:"code"`
					Type        string `json:"type"`
					Description string `json:"description"`
				} `json:"errors"`
			} `json:"search"`
		} `json:"hotelX"`
	} `json:"data"`
}

func NewResponse(options []*presenter.Option) Response {
	ret := Response{}
	ret.Data.HotelX.Search.Options = newOptions(options)
	return ret
}

type Option struct {
	SupplierCode      string        `json:"supplierCode"`
	AccessCode        string        `json:"accessCode"`
	Market            string        `json:"market"`
	HotelCode         string        `json:"hotelCode"`
	HotelCodeSupplier string        `json:"hotelCodeSupplier"`
	HotelName         *string       `json:"hotelName"`
	BoardCode         string        `json:"boardCode"`
	BoardCodeSupplier string        `json:"boardCodeSupplier"`
	PaymentType       string        `json:"paymentType"`
	Status            string        `json:"status"`
	Occupancies       []Occupancy   `json:"occupancies"`
	Rooms             []Room        `json:"rooms"`
	Price             Price         `json:"price"`
	Supplements       []Supplement  `json:"supplements"`
	Surcharges        []Surcharge   `json:"surcharges"`
	RateRules         []string      `json:"rateRules"`
	CancelPolicy      *CancelPolicy `json:"cancelPolicy"`
	Remarks           *string       `json:"remarks"`
	Token             string        `json:"token"`
	Id                string        `json:"id"`
}

type Surcharge struct {
	ChargeType  string  `json:"chargeType"`
	Mandatory   bool    `json:"mandatory"`
	Price       Price   `json:"price"`
	Description *string `json:"description"`
}

type CancelPolicy struct {
	Refundable      bool            `json:"refundable"`
	CancelPenalties []CancelPenalty `json:"cancelPenalties"`
}

type CancelPenalty struct {
	HoursBefore int     `json:"hoursBefore"`
	PenaltyType string  `json:"penaltyType"`
	Currency    string  `json:"currency"`
	Value       float64 `json:"value"`
}

type Supplement struct {
	Code           string  `json:"code"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	SupplementType string  `json:"supplementType"`
	ChargeType     string  `json:"chargeType"`
	Mandatory      bool    `json:"mandatory"`
	DurationType   *string `json:"durationType"`
	Quantity       *int    `json:"quantity"`
	Unit           *string `json:"unit"`
	EffectiveDate  *string `json:"effectiveDate"`
	ExpireDate     *string `json:"expireDate"`
	Resort         *Resort `json:"resort"`
	Price          *Price  `json:"price"`
}
type Resort struct {
	Code        string  `json:"code"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type Room struct {
	OccupancyRefId int         `json:"occupancyRefId"`
	Code           string      `json:"code"`
	Description    *string     `json:"description"`
	Refundable     *bool       `json:"refundable"`
	Units          *int        `json:"units"`
	RoomPrice      RoomPrice   `json:"roomPrice"`
	Beds           []Bed       `json:"beds"`
	RatePlans      []RatePlan  `json:"ratePlans"`
	Promotions     []Promotion `json:"promotions"`
}
type Promotion struct {
	Code          string  `json:"code"`
	Name          *string `json:"name"`
	EffectiveDate *string `json:"effectiveDate"`
	ExpireDate    *string `json:"expireDate"`
}

type RatePlan struct {
	Code          string  `json:"code"`
	Name          *string `json:"name"`
	EffectiveDate *string `json:"effectiveDate"`
	ExpireDate    *string `json:"expireDate"`
}

type Bed struct {
	Type        *string `json:"type"`
	Description *string `json:"description"`
	Count       *int    `json:"count"`
	Shared      *bool   `json:"shared"`
}

type RoomPrice struct {
	Price     Price            `json:"price"`
	Breakdown []PriceBreakdown `json:"breakdown"`
}

type PriceBreakdown struct {
	EffectiveDate string `json:"effectiveDate"`
	ExpireDate    string `json:"expireDate"`
	Price         Price  `json:"price"`
}
type Price struct {
	Currency string   `json:"currency"`
	Binding  bool     `json:"binding"`
	Net      float64  `json:"net"`
	Gross    *float64 `json:"gross"`
	Exchange Exchange `json:"exchange"`
	Markups  []Markup `json:"markups"`
}
type Markup struct {
	Channel  *string  `json:"channel"`
	Currency string   `json:"currency"`
	Binding  bool     `json:"binding"`
	Net      float64  `json:"net"`
	Gross    *float64 `json:"gross"`
	Exchange Exchange `json:"exchange"`
	Rules    []Rule   `json:"rules"`
}

type Rule struct {
	Id    string  `json:"id"`
	Name  *string `json:"name"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}

type Exchange struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

type Occupancy struct {
	Id    int   `json:"id"`
	Paxes []Pax `json:"paxes"`
}

type Pax struct {
	Age int `json:"age"`
}

func newOptions(options []*presenter.Option) []Option {
	ret := make([]Option, 0, len(options))
	for key := range options {
		ret = append(ret, newOption(options[key]))
	}
	return ret
}

func newOption(option *presenter.Option) Option {
	ret := Option{}
	ret.Id = option.OptionID
	ret.SupplierCode = option.Supplier
	ret.AccessCode = option.Access
	ret.Market = option.Market
	ret.HotelCode = option.HotelCode
	ret.HotelCodeSupplier = option.HotelCode
	ret.HotelName = option.HotelName
	ret.BoardCode = *option.BoardCode
	ret.BoardCodeSupplier = option.BoardCodeOriginal
	ret.PaymentType = string(option.PaymentType)
	ret.Status = string(option.Status)
	ret.Remarks = option.Remarks
	ret.Token = option.Token
	if len(option.Occupancies) != 0 {
		ret.Occupancies = make([]Occupancy, 0, len(option.Occupancies))
		for key := range option.Occupancies {
			ret.Occupancies = append(ret.Occupancies, newOccupancy(option.Occupancies[key]))
		}
	}

	if len(option.Rooms) != 0 {
		ret.Rooms = make([]Room, 0, len(option.Rooms))
		for key := range option.Rooms {
			ret.Rooms = append(ret.Rooms, newRoom(option.Rooms[key]))
		}
	}

	ret.Price = *newPrice(&option.Price)
	if len(option.Supplements) != 0 {
		ret.Supplements = make([]Supplement, 0, len(option.Supplements))
		for key := range option.Supplements {
			ret.Supplements = append(ret.Supplements, newSupplement(option.Supplements[key]))
		}
	}

	if len(option.Surcharges) != 0 {
		ret.Surcharges = make([]Surcharge, 0, len(option.Surcharges))
		for key := range option.Surcharges {
			ret.Surcharges = append(ret.Surcharges, newSurcharges(option.Surcharges[key]))
		}
	}

	if len(option.RateRules) != 0 {
		ret.RateRules = make([]string, 0, len(option.RateRules))
		for key := range option.RateRules {
			ret.RateRules = append(ret.RateRules, newRateRule(option.RateRules[key]))
		}
	}

	ret.CancelPolicy = newCancelPolicy(option.CancelPolicy)
	return ret
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
			ret.CancelPenalties = make([]CancelPenalty, 0, len(policy.CancelPenalties))
			for key := range policy.CancelPenalties {
				ret.CancelPenalties = append(ret.CancelPenalties, newCancelPenalty(policy.CancelPenalties[key]))
			}
		}
	}
	return &ret
}

func newCancelPenalty(penalty domainHotelCommon.CancelPenalty) CancelPenalty {
	ret := CancelPenalty{
		HoursBefore: penalty.HoursBefore,
		PenaltyType: string(penalty.Type),
		Currency:    penalty.Currency,
		Value:       penalty.Value,
	}
	return ret
}

func newRateRule(rulesType access.RateRulesType) string {
	return string(rulesType)
}

func newSurcharges(surcharge domainHotelCommon.Surcharge) Surcharge {
	ret := Surcharge{
		ChargeType:  string(surcharge.ChargeType),
		Mandatory:   surcharge.Mandatory,
		Description: surcharge.Description,
	}
	price := newPrice(&surcharge.Price)
	if price != nil {
		ret.Price = *price
	} else {
		ret.Price = Price{}
	}
	return ret
}

func newSupplement(supplement *domainHotelCommon.Supplement) Supplement {
	if supplement == nil {
		return Supplement{}
	}
	ret := Supplement{
		Name:           supplement.Name,
		Description:    supplement.Description,
		SupplementType: string(supplement.SupplementType),
		ChargeType:     string(supplement.ChargeType),
		Mandatory:      supplement.Mandatory,
	}
	if supplement.Code == nil {
		ret.Code = ""
	} else {
		ret.Code = *supplement.Code
	}
	if supplement.DurationType == nil {
		ret.DurationType = nil
	} else {
		s := string(*supplement.DurationType)
		ret.DurationType = &s
	}
	ret.Quantity = supplement.Quantity
	if supplement.Unit == nil {
		ret.Unit = nil
	} else {
		s := string(*supplement.Unit)
		ret.Unit = &s
	}
	ret.EffectiveDate = supplement.EffectiveDate
	ret.ExpireDate = supplement.ExpireDate
	ret.Resort = newResort(supplement.Resort)
	ret.Price = newPrice(supplement.Price)
	return ret
}

func newResort(resort *domainHotelCommon.Resort) *Resort {
	if resort == nil {
		return nil
	}
	ret := Resort{
		Code:        resort.Code,
		Name:        resort.Name,
		Description: resort.Description,
	}
	return &ret
}

func newOccupancy(occupancy domainHotelCommon.Occupancy) Occupancy {
	ret := Occupancy{
		Id: occupancy.Id,
	}

	if len(occupancy.Paxes) != 0 {
		ret.Paxes = make([]Pax, 0, len(occupancy.Paxes))
		for key := range occupancy.Paxes {
			ret.Paxes = append(ret.Paxes, newPax(occupancy.Paxes[key]))
		}
	}
	return ret
}

func newPax(pax domainHotelCommon.Pax) Pax {
	return Pax{
		pax.Age,
	}
}

func newRoom(room domainHotelCommon.Room) Room {
	ret := Room{
		OccupancyRefId: int(room.OccupancyRefID),
		Description:    room.Description,
		Refundable:     room.Refundable,
		Units:          room.Units,
	}
	if room.Code == nil {
		ret.Code = ""
	} else {
		ret.Code = *room.Code
	}
	ret.RoomPrice = newRoomPrice(room.RoomPrice)
	if len(room.Beds) != 0 {
		ret.Beds = make([]Bed, 0, len(room.Beds))
		for key := range room.Beds {
			ret.Beds = append(ret.Beds, newBed(room.Beds[key]))
		}
	}
	if len(room.RatePlans) != 0 {
		ret.RatePlans = make([]RatePlan, 0, len(room.RatePlans))
		for key := range room.RatePlans {
			ret.RatePlans = append(ret.RatePlans, newRatePlan(room.RatePlans[key]))
		}
	}
	if len(room.Promotions) != 0 {
		ret.Promotions = make([]Promotion, 0, len(room.Promotions))
		for key := range room.Promotions {
			ret.Promotions = append(ret.Promotions, newPromotion(room.Promotions[key]))
		}
	}
	return ret
}

func newPromotion(promotion domainHotelCommon.Promotion) Promotion {
	ret := Promotion{
		Code:          promotion.Code,
		Name:          promotion.Name,
		EffectiveDate: promotion.EffectiveDate,
		ExpireDate:    promotion.ExpireDate,
	}
	return ret
}

func newRatePlan(plan domainHotelCommon.RatePlan) RatePlan {
	ret := RatePlan{
		Name:          plan.Name,
		EffectiveDate: plan.EffectiveDate,
		ExpireDate:    plan.ExpireDate,
	}
	if plan.Code == nil {
		ret.Code = ""
	} else {
		ret.Code = *plan.Code
	}
	return ret
}

func newBed(bed domainHotelCommon.Bed) Bed {
	ret := Bed{
		Type:        bed.Type,
		Description: bed.Description,
		Count:       bed.Count,
		Shared:      bed.Shared,
	}
	return ret
}

func newRoomPrice(price domainHotelCommon.RoomPrice) RoomPrice {
	ret := RoomPrice{}
	ret.Price = *newPrice(&price.Price)
	if len(price.Breakdown) != 0 {
		ret.Breakdown = make([]PriceBreakdown, 0, len(price.Breakdown))
		for key := range price.Breakdown {
			ret.Breakdown = append(ret.Breakdown, newPriceBreakdown(price.Breakdown[key]))
		}
	}
	return ret
}

func newPriceBreakdown(down domainHotelCommon.PriceBreakDown) PriceBreakdown {
	ret := PriceBreakdown{
		EffectiveDate: down.EffectiveDate,
		ExpireDate:    down.ExpireDate,
		Price:         *newPrice(&down.Price),
	}
	return ret
}

func newPrice(price *domainHotelCommon.Price) *Price {
	if price == nil {
		return nil
	}
	ret := Price{
		Currency: price.Currency,
		Binding:  price.Binding,
		Net:      price.Net,
		Gross:    &price.Gross,
		Exchange: newExchange(price.Exchange),
	}
	if len(price.Markups) != 0 {
		ret.Markups = make([]Markup, 0, len(price.Markups))
		for key := range price.Markups {
			ret.Markups = append(ret.Markups, newMarkup(price.Markups[key]))
		}
	}
	return &ret
}

func newMarkup(markup domainHotelCommon.Markup) Markup {
	ret := Markup{
		Channel:  markup.Channel,
		Currency: markup.Currency,
		Binding:  markup.Binding,
		Net:      markup.Net,
		Gross:    &markup.Gross,
		Exchange: newExchange(markup.Exchange),
	}
	if len(markup.Rules) != 0 {
		ret.Rules = make([]Rule, 0, len(markup.Rules))
		for key := range markup.Rules {
			ret.Rules = append(ret.Rules, newRule(markup.Rules[key]))
		}
	}
	return ret
}

func newRule(rule domainHotelCommon.Rule) Rule {
	ret := Rule{
		Id:    rule.Id,
		Name:  rule.Name,
		Type:  string(rule.Type),
		Value: rule.Value,
	}
	return ret
}

func newExchange(exchange domainHotelCommon.Exchange) Exchange {
	ret := Exchange{
		Currency: exchange.Currency,
		Rate:     exchange.Rate,
	}
	return ret
}
