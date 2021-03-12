package main

import (
	"covid-19-graphql-service/pkg/schema"
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	env "github.com/joho/godotenv"
)

func init() {
	env.Load()
}

func main() {
	_, playground := os.LookupEnv("GRAPHQL_PLAYGROUND")

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Query,
	})

	handler := handler.New(&handler.Config{
		Schema:     &schema,
		Playground: playground,
	})

	fmt.Println("Starting COVID-19 GraphQL Service!...")
	fmt.Println("Server listening on http://localhost:4567/graphql")

	http.Handle("/graphql", handler)
	http.ListenAndServe(":4567", nil)
}
