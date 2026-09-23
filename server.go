package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/AKSHAYK0UL/graph-ql/config"
	"github.com/AKSHAYK0UL/graph-ql/graph"
)

func main() {

	if err := config.ConnectDB(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{}},
		),
	)

	http.Handle("/graphql", srv)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
