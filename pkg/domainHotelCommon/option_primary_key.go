package domainHotelCommon

import (
	"errors"
	"hub-aggregator/common/utils/conv"
	"sort"
	"strings"
)

func PrimaryKeyIndexFunc(pkChain string) (OptionIndexFunc, error) {
	indexFuncs, err := fromPkChainToIndexFuncs(pkChain)
	if err != nil {
		return nil, err
	}
	return func(opt *Option) string {
		// Si alguien se anima a hacer todos los casos... <3 (es mucho mas eficiente que el loop con el builder)
		switch len(indexFuncs) {
		case 1:
			return indexFuncs[0](opt)
		case 2:
			return indexFuncs[0](opt) + "#" + indexFuncs[1](opt)
		case 3:
			return indexFuncs[0](opt) + "#" + indexFuncs[1](opt) + "#" + indexFuncs[2](opt)
		case 4:
			return indexFuncs[0](opt) + "#" + indexFuncs[1](opt) + "#" + indexFuncs[2](opt) + "#" + indexFuncs[3](opt)
		}

		sb := strings.Builder{}
		for i, index := range indexFuncs {
			if i > 0 && i < len(indexFuncs)-1 {
				sb.WriteString("#")
			}
			sb.WriteString(index(opt))
		}
		return sb.String()
	}, nil
}

func fromPkChainToIndexFuncs(pkChain string) ([]OptionIndexFunc, error) {
	sptChain := strings.Split(pkChain, ",")
	indexFuncs := make([]OptionIndexFunc, len(sptChain))
	for i, k := range sptChain {
		if f, ok := pkChainMap[strings.Trim(k, " ")]; !ok {
			return nil, errPrimaryKey(k)
		} else {
			indexFuncs[i] = f
		}
	}
	return indexFuncs, nil
}

var pkChainMap = map[string]OptionIndexFunc{
	"currency":     OptionIndexCurrency,
	"supplier":     OptionIndexSupplier,
	"hotel":        OptionIndexHotelCode,
	"market":       OptionIndexMarket,
	"board":        OptionIndexBoard,
	"payment":      OptionIndexPayment,
	"room":         OptionIndexRoom,
	"promotion":    OptionIndexPromotion,
	"supplement":   OptionIndexSupplement,
	"surcharges":   OptionIndexSurcharges,
	"rateRules":    OptionIndexRateRules,
	"cancelPolicy": OptionIndexCancelPolicy,
}

func OptionIndexCurrency(opt *Option) string  { return opt.Price.Currency }
func OptionIndexSupplier(opt *Option) string  { return opt.Supplier }
func OptionIndexHotelCode(opt *Option) string { return opt.HotelCode }
func OptionIndexMarket(opt *Option) string    { return opt.Market }
func OptionIndexBoard(opt *Option) string     { return *opt.BoardCode }
func OptionIndexPayment(opt *Option) string   { return opt.PaymentType.Description() }

func OptionIndexRoom(opt *Option) string {
	roomCodes := make([]string, 0, len(opt.Rooms))
	for _, room := range opt.Rooms {
		roomCodes = append(roomCodes, *room.Code)
	}
	sort.Strings(roomCodes)
	return strings.Join(roomCodes, "|")
}

func OptionIndexPromotion(opt *Option) string {
	roomPromotions := make([]string, 0, len(opt.Rooms))
	for _, room := range opt.Rooms {
		if len(room.Promotions) != 0 {
			promotions := make([]string, 0, len(room.Promotions))
			for _, promotion := range room.Promotions {
				promotions = append(promotions, promotion.Code)
			}
			sort.Strings(promotions)
			roomPromotions = append(roomPromotions, *room.Code+"-"+strings.Join(promotions, "@"))
		}
	}
	sort.Strings(roomPromotions)
	return strings.Join(roomPromotions, "|")
}

func OptionIndexSupplement(opt *Option) string {
	if len(opt.Supplements) == 0 {
		return ""
	}

	supplements := make([]string, 0, len(opt.Supplements))
	for _, supplement := range opt.Supplements {
		supplements = append(supplements, *supplement.Code+"-"+supplement.SupplementType.String())
	}

	sort.Strings(supplements)
	return strings.Join(supplements, "|")
}
func OptionIndexSurcharges(opt *Option) string {
	if len(opt.Surcharges) == 0 {
		return ""
	}

	surcharges := make([]string, 0, len(opt.Surcharges))
	for _, surcharge := range opt.Surcharges {
		surcharges = append(surcharges, *surcharge.Description)
	}
	sort.Strings(surcharges)
	return strings.Join(surcharges, "|")
}

func OptionIndexRateRules(opt *Option) string {
	if len(opt.RateRules) == 0 {
		return ""
	}

	rateRules := make([]string, 0, len(opt.RateRules))
	for _, rr := range opt.RateRules {
		rateRules = append(rateRules, rr.String())
	}
	sort.Strings(rateRules)
	return strings.Join(rateRules, "|")

}

func OptionIndexCancelPolicy(opt *Option) string {
	if opt.CancelPolicy == nil {
		return "0" // no refundable
	}
	if len(opt.CancelPolicy.CancelPenalties) == 0 {
		return conv.BoolToBitString(opt.CancelPolicy.Refundable)
	}

	cancelPolicies := make([]string, 0, len(opt.CancelPolicy.CancelPenalties))
	for _, cancelPenalty := range opt.CancelPolicy.CancelPenalties {
		cancelPolicies = append(cancelPolicies, cancelPenalty.Currency+"@"+
			cancelPenalty.Type.String()+"@"+conv.Itoa(cancelPenalty.HoursBefore)+"@"+
			conv.FormatFloat64(cancelPenalty.Value))
	}

	sort.Strings(cancelPolicies)
	return conv.BoolToBitString(opt.CancelPolicy.Refundable) + strings.Join(cancelPolicies, "|")
}

func errPrimaryKey(invalidValue string) error {
	return errors.New("wrong value found in PrimaryKey chain: " + invalidValue + ", use these fields: " +
		"supplier, hotel, market, board, payment, room, promotion, supplement, surcharges, rateRules or/and cancelPolicy splitted by ','")
}
