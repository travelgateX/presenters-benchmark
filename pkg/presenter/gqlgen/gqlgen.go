package gqlgen

import (
	"github.com/99designs/gqlgen/handler"
	"net/http"
	"github.com/travelgateX/presenters-benchmark/pkg/presenter"
)

type Candidate struct{}

var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	return handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{options}})), nil
}

func (Candidate) UnmarshalOptions(b []byte) ([]*presenter.Option, error) {
	return presenter.JSONUnmarshalOptions(b)
}
