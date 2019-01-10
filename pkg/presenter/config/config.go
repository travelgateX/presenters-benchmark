package config

import (
	"hub-aggregator/common/config/log"
	"hub-aggregator/common/kit/routing"
	"rfc/presenters/pkg/presenter"
	"rfc/presenters/pkg/presenter/gophers"
	"rfc/presenters/pkg/presenter/gqlgen"
	"rfc/presenters/pkg/presenter/rest"
)

type Config struct {
	Logger log.Builder
}

func NewRoutes(logger log.Builder) []routing.Route {

	restService := presenter.NewService(
		presenter.NewOptionsGen(),
		presenter.NewSearchGraphQLRequester(),
		presenter.NewLogger(logger.Logger()),
		rest.Candidate{},
	)
	gqlgenService := presenter.NewService(
		presenter.NewOptionsGen(),
		presenter.NewSearchGraphQLRequester(),
		presenter.NewLogger(logger.Logger()),
		gqlgen.Candidate{},
	)
	gophersService := presenter.NewService(
		presenter.NewOptionsGen(),
		presenter.NewSearchGraphQLRequester(),
		presenter.NewLogger(logger.Logger()),
		gophers.Candidate{},
	)

	return []routing.Route{
		routing.Route{
			Name:              "expgraph rest",
			Method:            "POST",
			Pattern:           "/exp/graph/rest",
			ServiceHandleFunc: restService.HandlerFunc(),
		},
		routing.Route{
			Name:              "expgraph gqlgen",
			Method:            "POST",
			Pattern:           "/exp/graph/gqlgen",
			ServiceHandleFunc: gqlgenService.HandlerFunc(),
		},
		routing.Route{
			Name:              "expgraph gophers",
			Method:            "POST",
			Pattern:           "/exp/graph/gophers",
			ServiceHandleFunc: gophersService.HandlerFunc(),
		},
	}
}
