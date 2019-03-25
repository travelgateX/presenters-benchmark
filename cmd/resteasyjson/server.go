package main

import (
	"github.com/99designs/gqlgen/handler"
	"log"
	"net/http"
	"os"
	"presenters-benchmark/pkg/presenter"
	"presenters-benchmark/pkg/presenter/resteasyjson"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	options := presenter.NewOptionsGen().Gen(1)
	h, err := resteasyjson.Candidate{}.HandlerFunc(options)
	if err != nil {
		panic(err)
	}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", h)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
