package domainHotelCommon

import (
	"github.com/travelgateX/presenters-benchmark/pkg/access"
)

type OptionID struct {
	StartDate     string
	EndDate       string
	HotelCode     string
	BoardCode     string
	PaymentType   PaymentType
	Market        string
	Nationality   string
	Language      string
	Currency      string
	OptionsQuota  *int32
	BusinessRules *BusinessRulesType
	AccessCode    string
	Rooms         []OptionIdRoom
	Parameters    []access.Parameter `json:"p"`
	Price         *OptionIdPrice
	// markup needed fields, we need codes in the context request in search
	ContextHC string
	// needs to be a pointer, its value is updated after parsing response, since OptionId is generated at resolving time
	// it will contain the correct value
	ContextBoard *string
	Rate         Rate
	// only from quote to book
	Refundable bool
}

type OptionIdRoom struct {
	RefID int32
	ID    string
	Code  string
	Paxes []Pax
}

type OptionIdPrice struct {
	Value                   string
	CommissionIntegration   string
	Binding                 string
	Currency                string
	CurrencyEx              string
	RateEx                  float64
	CommissionGross         float64
	CommissionNet           float64
	CommissionPluginApplied bool
	IsOldToken              bool
}
