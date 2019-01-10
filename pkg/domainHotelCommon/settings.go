package domainHotelCommon

type Settings struct {
	SettingsBase

	Context      *string
	Client       *string
	Group        *string
	Org          string
	Suppliers    *[]Supplier
	Plugins      *[]PluginStep
	TestMode     *bool
	ClientTokens *[]string

	// deprecated
	UseContext  *bool
	ConnectUser *string

	// pinxo
	IsClientNewCommission_pinxo bool
	IsClientMap_pinxo           bool
	IsAvorisClient_pinxo        bool

	IsClientDiscardBoardMap_pinxo bool
	IsClientDiscardRoomMap_pinxo  bool
	IsClientDiscardRateMap_pinxo  bool
}

type Plugins struct {
	Plugins *[]PluginStep
	Filter  *PluginFilter
}
type PluginFilter struct {
	Plugin *FilterPluginType
}
