package main

import (
	"log"
	"net/http"
	"os"
	"rfc/presenters/pkg/domainHotelCommon"
	"rfc/presenters/pkg/presenter"
	"rfc/presenters/pkg/presenter/gqlgensm"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	options := presenter.NewOptionsGen().Gen(1)
	soptions := make([]*domainHotelCommon.Option, len(options))
	for i, o := range options {
		opt := (domainHotelCommon.Option)(*o)
		soptions[i] = &opt
	}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &gqlgen.Resolver{soptions}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
