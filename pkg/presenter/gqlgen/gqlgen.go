package gqlgen

import (
	"github.com/99designs/gqlgen/handler"
	"hub-aggregator/common/kit/routing"
	"net/http"
	"rfc/presenters/pkg/presenter"
)

type Candidate struct{}

var _ presenter.CandidateServer = (*Candidate)(nil)
var _ presenter.CandidateHandlerFunc = (*Candidate)(nil)

func (Candidate) NewServer(addr, pattern string, options []*presenter.Option, results chan<- presenter.OperationResult) (*routing.Server, error) {
	return presenter.NewGzipCandidateServer(
		addr,
		pattern,
		handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{options}})),
		presenter.TotalTimeMiddleware(results),
	), nil
}

func (Candidate) HandlerFunc(options []*presenter.Option) (http.HandlerFunc, error) {
	return handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{options}})), nil
}
