package domainHotelCommon

type Price struct {
	Currency                string   `json:"currency"`
	Binding                 bool     `json:"binding"`
	Net                     float64  `json:"net"`
	Gross                   float64  `json:"gross"`
	CommissionIntegration   float64  `json:"-"`
	CommissionPluginApplied bool     `json:"-"`
	Exchange                Exchange `json:"exchange"`
	Markups                 []Markup `json:"markups"`
	CommissionNet           float64  `json:"-"`
	CommissionGross         float64  `json:"-"`
}

//don't use if NewPrice is not used previously
func (p Price) ApplyCommissionBlue(comm float64) Price {
	p.Net = p.Gross * (1.0 - (p.CommissionIntegration / 100.0))
	p.Gross = p.Net * (1.0 + comm/100.0)
	p.CommissionGross = 100.0 - (100.0 * p.Net / p.Gross)
	p.CommissionNet = comm
	p.CommissionPluginApplied = true
	return p
}

//don't use if NewPrice is not used previously
func (p Price) ApplyCommissionRed(comm float64) Price {
	p.Net = p.Gross * (1.0 - (comm / 100.0))
	p.CommissionNet = (100.0 * p.Gross / p.Net) - 100.0
	p.CommissionGross = comm
	p.CommissionPluginApplied = true
	return p
}

func (Price) IsPriceable() {}
