package domainHotelCommon

import (
	"fmt"
	"io"
	"presenters-benchmark/pkg/access"
	"strconv"
)

type PluginStep struct {
	Step        PluginStepType
	PluginsType *[]Plugin
}

type Plugin struct {
	Name       string
	Type       PluginType
	Parameters *[]access.Parameter
}

type PluginStepType string

const (
	PluginStepTypeRequest        PluginStepType = "REQUEST"
	PluginStepTypeRequestAccess  PluginStepType = "REQUEST_ACCESS"
	PluginStepTypeResponseOption PluginStepType = "RESPONSE_OPTION"
	PluginStepTypeResponseAccess PluginStepType = "RESPONSE_ACCESS"
	PluginStepTypeResponse       PluginStepType = "RESPONSE"
)

func (e PluginStepType) IsValid() bool {
	switch e {
	case PluginStepTypeRequest, PluginStepTypeRequestAccess, PluginStepTypeResponseOption, PluginStepTypeResponseAccess, PluginStepTypeResponse:
		return true
	}
	return false
}

func (e PluginStepType) String() string {
	return string(e)
}

func (e *PluginStepType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PluginStepType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PluginStepType", str)
	}
	return nil
}

func (e PluginStepType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PluginType string

const (
	PluginTypePreStep            PluginType = "PRE_STEP"
	PluginTypeHotelMap           PluginType = "HOTEL_MAP"
	PluginTypeBoardMap           PluginType = "BOARD_MAP"
	PluginTypeRoomMap            PluginType = "ROOM_MAP"
	PluginTypeCurrencyConversion PluginType = "CURRENCY_CONVERSION"
	PluginTypeMarkup             PluginType = "MARKUP"
	PluginTypeAggregation        PluginType = "AGGREGATION"
	PluginTypePostStep           PluginType = "POST_STEP"
)

func (e PluginType) IsValid() bool {
	switch e {
	case PluginTypePreStep, PluginTypeHotelMap, PluginTypeBoardMap, PluginTypeRoomMap, PluginTypeCurrencyConversion, PluginTypeMarkup, PluginTypeAggregation, PluginTypePostStep:
		return true
	}
	return false
}

func (e PluginType) String() string {
	return string(e)
}

func (e *PluginType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PluginType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PluginType", str)
	}
	return nil
}

func (e PluginType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterPluginType struct {
	Includes *[]FilterPlugin
	Excludes *[]FilterPlugin
}

type FilterPlugin struct {
	Step PluginStepType
	Type PluginType
	Name string
}
