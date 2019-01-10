package access

import (
	"fmt"
	"io"
	"strconv"
)

type AccessConfiguration struct {
	Code       string       `json:"code"`
	IsActive   int32        `json:"act"`
	IsTest     int32        `json:"isTest"`
	Supplier   string       `json:"hubProvider"`
	Username   *string      `json:"user,omitempty"`
	Password   *string      `json:"password,omitempty"`
	Urls       Urls         `json:"urls,omitempty"`
	Parameters *[]Parameter `json:"parameters,omitempty"`
	// Markets to which this access is valid
	Markets   *[]string        `json:"markets,omitempty"`
	RateRules *[]RateRulesType `json:"rateRules,omitempty"`
}

type RateRulesType string

const (
	RateRulesTypePackage          RateRulesType = "PACKAGE"
	RateRulesTypeOlder55          RateRulesType = "OLDER55"
	RateRulesTypeOlder60          RateRulesType = "OLDER60"
	RateRulesTypeOlder65          RateRulesType = "OLDER65"
	RateRulesTypeCanaryResident   RateRulesType = "CANARY_RESIDENT"
	RateRulesTypeBalearicResident RateRulesType = "BALEARIC_RESIDENT"
	RateRulesTypeLargeFamily      RateRulesType = "LARGE_FAMILY"
	RateRulesTypeHoneymoon        RateRulesType = "HONEYMOON"
	RateRulesTypePublicServant    RateRulesType = "PUBLIC_SERVANT"
	RateRulesTypeUnemployed       RateRulesType = "UNEMPLOYED"
	RateRulesTypeNormal           RateRulesType = "NORMAL"
	RateRulesTypeNonRefundable    RateRulesType = "NON_REFUNDABLE"
)

func (e RateRulesType) IsValid() bool {
	switch e {
	case RateRulesTypePackage, RateRulesTypeOlder55, RateRulesTypeOlder60, RateRulesTypeOlder65, RateRulesTypeCanaryResident, RateRulesTypeBalearicResident, RateRulesTypeLargeFamily, RateRulesTypeHoneymoon, RateRulesTypePublicServant, RateRulesTypeUnemployed, RateRulesTypeNormal, RateRulesTypeNonRefundable:
		return true
	}
	return false
}

func (e RateRulesType) String() string {
	return string(e)
}

func (e *RateRulesType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RateRulesType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RateRulesType", str)
	}
	return nil
}

func (e RateRulesType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (a *AccessConfiguration) Clone() *AccessConfiguration {
	v := *a

	if a.Username != nil {
		v.Username = new(string)
		*v.Username = *a.Username
	}

	if a.Password != nil {
		v.Password = new(string)
		*v.Password = *a.Password
	}

	if a.Parameters != nil {
		v.Parameters = new([]Parameter)
		*v.Parameters = make([]Parameter, len(*a.Parameters))
		copy(*v.Parameters, *a.Parameters)
	}

	if a.Markets != nil {
		v.Markets = new([]string)
		*v.Markets = make([]string, len(*a.Markets))
		copy(*v.Markets, *a.Markets)
	}

	if a.RateRules != nil {
		v.RateRules = new([]RateRulesType)
		*v.RateRules = make([]RateRulesType, len(*a.RateRules))
		copy(*v.RateRules, *a.RateRules)
	}

	return &v
}

type Urls struct {
	Search  *string `json:"search,omitempty"`
	Quote   *string `json:"quote,omitempty"`
	Book    *string `json:"book,omitempty"`
	Generic *string `json:"generic,omitempty"`
}
