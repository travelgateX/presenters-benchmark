package presenter

import (
	"encoding/json"
	"reflect"
	"presenters-benchmark/pkg/domainHotelCommon"
)

type Option domainHotelCommon.Option

func OptionEquals(opt1, opt2 *Option) bool {
	return reflect.DeepEqual(opt1, opt2)
}

type OptionsGen interface {
	Gen(n int) []*Option
}

func NewOptionsGen() OptionsGen {
	return optionsGen{}
}

type optionsGen struct{}

func (optionsGen) Gen(n int) []*Option {
	var optObj domainHotelCommon.Option
	err := json.Unmarshal(mockOption(), &optObj)
	if err != nil {
		return nil
	}

	ret := make([]*Option, 0, n)
	for k := 0; k < n; k++ {
		e := Option(optObj)
		ret = append(ret, &e)
	}

	return ret
}

func mockOption() []byte {
	return []byte(`{
            "id": "41[01|190415|190421|75040|TI|0|GB|GB|en|GBP|0|1|829|56061|6|1|0|1#931|931|30#30|agencyToken[CgU3NTA0MBICVEkaCjE1LzA0LzIwMTkiCjIxLzA0LzIwMTkqC01lcmNoYW50UGF5MgNHQlA6EgoDR0JQEgY1OTMuOTYiAzAuMEIVCgUxIzAkzMRABGgM5MzEiBTMwLTMwSg8KATESBAgeEAESBAgeEAJSQQoIaWRPcGNpb24SNTc1MDQwfDIwMTkwNDE1fDZ8MzAtMzB8VUt8VEl8SHxZT1V8MHx8fHxNfEh8Nzg0fDkzMXx8Ug0vKCGNvZENhY2hlEgEwWgYwLjg5MjNiFjIwMTgtMTItMDRUMDE6MDU6MDIuODFoAXIECEIYAnp3ChUKAzk2MBIOCgNHQlAQAhoFNTkuNDAKFgoDNTc2Eg8KA0dCUBACGgYxNEDguNDkKFgoDMzYwEg8KA0dCUBACGgYyOTYuOTgKFgoDMjQwEg8KA0dCUBACGgY0NDUuNDcKFgoDMTIwEg8KA0dCUBACGgY1OTMuOTaCAQNZT1WSAQJHQsIBATDSAQEy2gEHUD0jTD1HQuIBBUFWQUlM8gEDNi42+gETMjAxOC0xMi0wNCAwNzo1NzozNYoCDjw4MzEzIyArMSBbMV0+kgIDMy42mgIDMi4wogIBMMoCATD6AgMwLjCCAwMwLjCKAwEw",
            "token": "VFJVfEhvdGVsWF8xNTV8ODI5fDU5My45NiM1OTMuOTYjMCNHQlB8MXwxOTA0MTV8MTkwNDIxfDAjMzAjMzB8ZW58R0JQfEdCfEdCfDU2MDYx",
            "accessCode": "829",
            "supplierCode": "MGR",
            "market": "GB",
            "hotelCode": "56061",
            "hotelCodeSupplier": "75040",
            "hotelName": null,
            "boardCode": "6",
            "boardCodeSupplier": "TI",
            "paymentType": "MERCHANT",
            "status": "OK",
            "addOns": null,
            "occupancies": [
              {
                "id": 1,
                "paxes": [
                  {
                    "age": 30
                  },
                  {
                    "age": 30
                  }
                ]
              }
            ],
            "rooms": [
              {
                "occupancyRefId": 1,
                "code": "apt sea side",
                "description": "Apartment - Lateral Sea View",
                "refundable": null,
                "units": null,
                "roomPrice": {
                  "price": {
                    "currency": "GBP",
                    "binding": false,
                    "net": 593.96,
                    "gross": 593.96,
                    "exchange": {
                      "currency": "GBP",
                      "rate": 1
                    },
                    "markups": null
                  }
                },
                "beds": null,
                "ratePlans": null,
                "promotions": null
              }
            ],
            "price": {
              "currency": "GBP",
              "binding": false,
              "net": 593.96,
              "gross": 593.96,
              "exchange": {
                "currency": "GBP",
                "rate": 1
              },
              "markups": null
            },
            "supplements": null,
            "surcharges": null,
            "rateRules": null,
            "cancelPolicy": {
              "refundable": true,
              "cancelPenalties": [
                {
                  "hoursBefore": 960,
                  "penaltyType": "IMPORT",
                  "currency": "GBP",
                  "value": 59.4
                },
                {
                  "hoursBefore": 576,
                  "penaltyType": "IMPORT",
                  "currency": "GBP",
                  "value": 148.49
                },
                {
                  "hoursBefore": 360,
                  "penaltyType": "IMPORT",
                  "currency": "GBP",
                  "value": 296.98
                },
                {
                  "hoursBefore": 240,
                  "penaltyType": "IMPORT",
                  "currency": "GBP",
                  "value": 445.47
                },
                {
                  "hoursBefore": 120,
                  "penaltyType": "IMPORT",
                  "currency": "GBP",
                  "value": 593.96
                }
              ]
            },
            "remarks": null
          }`)
}
