package graphResolver

import (
	"fmt"
	"github.com/travelgateX/presenters-benchmark/pkg/search"
	"strconv"
	"strings"
)

type CriteriaRequestResolver struct {
	Criteria *search.CriteriaSearch
}

func (r *CriteriaRequestResolver) CheckIn() Date {
	return Date(r.Criteria.CheckIn)
}

func (r *CriteriaRequestResolver) CheckOut() Date {
	return Date(r.Criteria.CheckOut)
}

func (r *CriteriaRequestResolver) Nationality() *Country {
	if r.Criteria.Nationality == nil {
		return nil
	}
	tmp := Country(*r.Criteria.Nationality)
	return &tmp
}

func (r *CriteriaRequestResolver) Currency() *Currency {
	if r.Criteria.Currency == nil {
		return nil
	}
	tmp := Currency(*r.Criteria.Currency)
	return &tmp
}

func (r *CriteriaRequestResolver) Language() *Language {
	if r.Criteria.Language == nil {
		return nil
	}
	tmp := Language(*r.Criteria.Language)
	return &tmp
}

func (r *CriteriaRequestResolver) Market() string {
	return *r.Criteria.Market
}

func (r *CriteriaRequestResolver) Hotels() []string {
	if r.Criteria.Hotels == nil {
		return []string{}
	}
	return *r.Criteria.Hotels
}

func (r *CriteriaRequestResolver) Occupancies() []*OccupancyResolver {
	occupancies := make([]*OccupancyResolver, 0, len(r.Criteria.Occupancies))
	for _, occupancy := range r.Criteria.Occupancies {
		occupancy_aux := occupancy
		occupancies = append(occupancies, &OccupancyResolver{Occupancy: &occupancy_aux})
	}
	return occupancies
}

type Date string

func (_ Date) ImplementsGraphQLType(name string) bool {
	return name == "Date"
}

func (date *Date) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*date = Date(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (date Date) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(date)), nil
}

type Language string

func (_ Language) ImplementsGraphQLType(name string) bool {
	return name == "Language"
}

func (language *Language) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*language = Language(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (language Language) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(language)), nil
}

// Country
type Country string

func (_ Country) ImplementsGraphQLType(name string) bool {
	return name == "Country"
}

func (country *Country) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*country = Country(strings.ToUpper(input))
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (country Country) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(country)), nil
}

// URI
type Uri string

func (_ Uri) ImplementsGraphQLType(name string) bool {
	return name == "URI"
}

func (uri *Uri) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*uri = Uri(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (uri Uri) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(uri)), nil
}

// CVC
type CVC string

func (CVC) ImplementsGraphQLType(name string) bool {
	return name == "CVC"
}

func (cvc *CVC) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*cvc = CVC(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (cvc CVC) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(cvc)), nil
}

// CardNumber
type CardNumber string

func (CardNumber) ImplementsGraphQLType(name string) bool {
	return name == "CardNumber"
}

func (cardNumber *CardNumber) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*cardNumber = CardNumber(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (cardNumber CardNumber) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(cardNumber)), nil
}
