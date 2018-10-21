package main

import (
	graphql_go "graphql-go"
	log "log"
	http "net/http"
	os "os"

	handler "github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql_go.NewExecutableSchema(graphql_go.Config{Resolvers: &graphql_go.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
