package main

import (
	"net/http"
	"simvino/auth"
	"simvino/config"
	"simvino/db"
	"simvino/graph"
	"simvino/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func main() {

	config.GetSecret()
	db.Init()
	router := chi.NewRouter()

	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":8070", router)
	if err != nil {
		panic(err)
	}
}
