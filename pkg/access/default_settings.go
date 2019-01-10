package access

type DefaultSettingsRedis struct {
	ConnectUser   string        `json:"hubUser"`
	Context       *string       `json:"context"`
	Language      string        `json:"language"`
	Currency      string        `json:"currency"`
	Nationality   string        `json:"nationality"`
	Market        string        `json:"market"`
	Timeout       Timeout       `json:"timeouts"`
	BusinessRules BusinessRules `json:"businessRules"`
}

type DefaultSettings struct {
	connectUser   string
	context       *string
	language      string
	currency      string
	nationality   string
	market        string
	timeout       Timeout
	businessRules BusinessRules
}

type BusinessRules struct {
	OptionsQuota      int32  `json:"optionsQuota"`
	BusinessRulesType string `json:"businessRulesType"`
}

type Timeout struct {
	Search int32 `json:"search"`
	Quote  int32 `json:"quote"`
	Book   int32 `json:"book"`
}

func NewDefaultSettings(dsr *DefaultSettingsRedis) *DefaultSettings {
	return &DefaultSettings{
		connectUser:   dsr.ConnectUser,
		context:       dsr.Context,
		language:      dsr.Language,
		currency:      dsr.Currency,
		nationality:   dsr.Nationality,
		timeout:       dsr.Timeout,
		businessRules: dsr.BusinessRules,
		market:        dsr.Market,
	}
}

func (ds *DefaultSettings) ConnectUser() string {
	return ds.connectUser
}

func (ds *DefaultSettings) Context() *string {
	return ds.context
}

func (ds *DefaultSettings) Language() string {
	return ds.language
}

func (ds *DefaultSettings) Currency() string {
	return ds.currency
}

func (ds *DefaultSettings) Nationality() string {
	return ds.nationality
}

func (ds *DefaultSettings) Market() string {
	return ds.market
}

func (ds *DefaultSettings) SearchTimeout() int32 {
	return ds.timeout.Search
}

func (ds *DefaultSettings) QuoteTimeout() int32 {
	return ds.timeout.Quote
}

func (ds *DefaultSettings) BookTimeout() int32 {
	return ds.timeout.Book
}

func (ds *DefaultSettings) OptionsQuota() int32 {
	return ds.businessRules.OptionsQuota
}

func (ds *DefaultSettings) BusinessRulesType() string {
	return ds.businessRules.BusinessRulesType
}
