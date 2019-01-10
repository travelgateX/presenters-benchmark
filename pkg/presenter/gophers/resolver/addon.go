package graphResolver

import (
	"hub-aggregator/common/graphql"
)

type AddOnResolver struct {
	key   string
	value string
}

func (r *AddOnResolver) Key() string {
	return r.key
}

func (r *AddOnResolver) Value() graphql.Json {
	return graphql.Json(r.value)
}
