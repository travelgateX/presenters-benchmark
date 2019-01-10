package presenter

type SearchGraphQLRequester interface {
	SearchGraphQLRequest(ResolveScale) []byte
}

// ResolveScale of fields in the graph Request
// low: options with 1 field
// medium: options with half the fields
// high: options with all the fields
type ResolveScale string

const (
	ResolveScaleLow    ResolveScale = "low"
	ResolveScaleMedium ResolveScale = "medium"
	ResolveScaleHigh   ResolveScale = "high"
)

type searchGraphQLRequester struct {
}

func NewSearchGraphQLRequester() SearchGraphQLRequester {
	return searchGraphQLRequester{}
}

var _ SearchGraphQLRequester = (*searchGraphQLRequester)(nil)

func (gr searchGraphQLRequester) SearchGraphQLRequest(level ResolveScale) []byte {
	switch level {
	case ResolveScaleLow:
		return gr.genRq(lowOptRq)
	case ResolveScaleMedium:
		return gr.genRq(medOptRq)
	case ResolveScaleHigh:
		return gr.genRq(highOptRq)
	}
	return nil
}

func (searchGraphQLRequester) genRq(optQuery string) []byte {
	return []byte(firstQueryPart + optQuery + lastQueryPart)
}

var firstQueryPart = `{"query":"query {\n  hotelX {\n    search`
var lastQueryPart = `"variables":{"criteriaSearch":{"market":"GB","nationality":"GB","checkOut":"2019-05-08","language":"en","checkIn":"2019-05-01","currency":"GBP","occupancies":[{"paxes":[{"age":30},{"age":30}]}],"hotels":["20698","43575","21545","46603","398238","45992","338172","398234","17043","349065","46609","297531","390513","45758","390515","390514","53385","53143","390516","30025","54479","2786","1454","53146","290904","55329","341337","250208","398249","45502","375587","398244","19231","398246","40297","405082","258188","1470","43329","34613","13936","400868","52066","54485","55574","31343","53152","35700","402803","53153","52064","55337","328360","32670","327274","3883","31109","402809","398250","89593","45734","43798","3660","398213","45730","40044","55340","54490","54259","13908","53168","6920","813","42211","17036","399318","399553","16183","54266","55597","10887","54265","54264","55594","29252","1001","59952","307420","265742","44864","358818","285754","2111","45717","400441","53184","53181","27285","31311","54277"]},"filter":null,"settings":{"auditTransactions":false,"client":"hoco_txt","clientTokens":["350"],"context":null,"plugins":[{"pluginsType":[{"name":"markup","parameters":[{"key":"channel","value":"350"}],"type":"MARKUP"}],"step":"RESPONSE_OPTION"},{"pluginsType":[{"name":"room_description_mapX","parameters":null,"type":"ROOM_MAP"}],"step":"RESPONSE_OPTION"},{"pluginsType":[{"name":"currency_exchange","parameters":[{"key":"currency","value":"GBP"},{"key":"exclude","value":"true"}],"type":"CURRENCY_CONVERSION"}],"step":"RESPONSE_OPTION"}],"testMode":false,"timeout":15000}}}`
var lowOptRq = `{\n      options {\n        id\n      }\n    }\n  }\n}\n",`
var medOptRq = `{\n      options {\n        price {\n          currency\n          binding\n          net\n          gross\n          exchange {\n            currency\n            rate\n          }\n          markups {\n            channel\n            currency\n            binding\n            net\n            gross\n            exchange {\n              currency\n              rate\n            }\n          }\n        }\n        addOns {\n          distribute\n        }\n        supplements {\n          code\n          name\n          description\n          supplementType\n          chargeType\n          mandatory\n          durationType\n          quantity\n          unit\n          effectiveDate\n          expireDate\n          resort {\n            code\n            name\n            description\n          }\n          price {\n            currency\n            binding\n            net\n            gross\n            exchange {\n              currency\n              rate\n            }\n            markups {\n              channel\n              currency\n              binding\n              net\n              gross\n              exchange {\n                currency\n                rate\n              }\n            }\n          }\n        }\n        surcharges {\n          chargeType\n          description\n          price {\n            currency\n            binding\n            net\n            gross\n            exchange {\n              currency\n              rate\n            }\n            markups {\n              channel\n              currency\n              binding\n              net\n              gross\n              exchange {\n                currency\n                rate\n              }\n            }\n          }\n        }\n        rateRules\n        cancelPolicy {\n          refundable\n          cancelPenalties {\n            hoursBefore\n            penaltyType\n            currency\n            value\n          }\n        }\n        remarks\n        token\n        id\n      }\n    }\n  }\n}\n",`
var highOptRq = `{\n      options {\n        hotelCode\n        surcharges {\n          chargeType\n          mandatory\n          description\n          price {\n            currency\n            binding\n            net\n            gross\n            exchange {\n              currency\n              rate\n            }\n            markups {\n              channel\n              currency\n              binding\n              net\n              gross\n              exchange {\n                currency\n                rate\n              }\n            }\n          }\n        }\n        accessCode\n        supplierCode\n        market\n        hotelCode\n        hotelName\n        boardCode\n        paymentType\n        status\n        occupancies {\n          id\n          paxes {\n            age\n          }\n        }\n        rooms {\n          occupancyRefId\n          code\n          description\n          refundable\n          units\n          roomPrice {\n            price {\n              currency\n              binding\n              net\n              gross\n              exchange {\n                currency\n                rate\n              }\n              markups {\n                channel\n                currency\n                binding\n                net\n                gross\n                exchange {\n                  currency\n                  rate\n                }\n              }\n            }\n          }\n          beds {\n            type\n            description\n            count\n            shared\n          }\n          ratePlans {\n            code\n            name\n            effectiveDate\n            expireDate\n          }\n          promotions {\n            code\n            name\n            effectiveDate\n            expireDate\n          }\n        }\n        price {\n          currency\n          binding\n          net\n          gross\n          exchange {\n            currency\n            rate\n          }\n          markups {\n            channel\n            currency\n            binding\n            net\n            gross\n            exchange {\n              currency\n              rate\n            }\n          }\n        }\n        addOns {\n          distribute\n        }\n        supplements {\n          code\n          name\n          description\n          supplementType\n          chargeType\n          mandatory\n          durationType\n          quantity\n          unit\n          effectiveDate\n          expireDate\n          resort {\n            code\n            name\n            description\n          }\n          price {\n            currency\n            binding\n            net\n            gross\n            exchange {\n              currency\n              rate\n            }\n            markups {\n              channel\n              currency\n              binding\n              net\n              gross\n              exchange {\n                currency\n                rate\n              }\n            }\n          }\n        }\n        surcharges {\n          chargeType\n          description\n          price {\n            currency\n            binding\n            net\n            gross\n            exchange {\n              currency\n              rate\n            }\n            markups {\n              channel\n              currency\n              binding\n              net\n              gross\n              exchange {\n                currency\n                rate\n              }\n            }\n          }\n        }\n        rateRules\n        cancelPolicy {\n          refundable\n          cancelPenalties {\n            hoursBefore\n            penaltyType\n            currency\n            value\n          }\n        }\n        remarks\n        token\n        id\n      }\n    }\n  }\n}\n",`
