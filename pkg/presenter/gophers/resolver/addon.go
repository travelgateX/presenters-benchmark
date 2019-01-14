package graphResolver

import (
	"fmt"
	"strconv"
)

type AddOnResolver struct {
	key   string
	value string
}

func (r *AddOnResolver) Key() string {
	return r.key
}

func (r *AddOnResolver) Value() Json {
	return Json(r.value)
}

type Json string

func (Json) ImplementsGraphQLType(name string) bool {
	return name == "JSON"
}

func (json *Json) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		*json = Json(input)
		return nil
	default:
		return fmt.Errorf("wrong type")
	}
}

func (json Json) MarshalJSON() ([]byte, error) {
	return strconv.AppendQuote(nil, string(json)), nil
}
