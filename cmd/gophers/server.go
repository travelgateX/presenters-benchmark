package main

import (
	"log"
	"net/http"
	"os"
	"rfc/presenters/pkg/domainHotelCommon"
	"rfc/presenters/pkg/presenter"
	"rfc/presenters/pkg/presenter/gophers"

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

	h, err := gophers.Candidate{}.HandlerFunc(options)
	if err != nil {
		panic(err)
	}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", h)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
