package domainHotelCommon

type Markup struct {
	Channel  *string  `json:"channel"`
	Currency string   `json:"currency"`
	Binding  bool     `json:"binding"`
	Net      float64  `json:"net"`
	Gross    float64  `json:"gross"`
	Exchange Exchange `json:"exchange"`
	Rules    []Rule   `json:"rules"`
}

func (Markup) IsPriceable() {}
