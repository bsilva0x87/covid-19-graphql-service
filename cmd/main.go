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
	err := env.Load()

	if err != nil {
		fmt.Println("Please add a .env file at the project root.\nCheck the README file.")
		os.Exit(0)
	}
}

func main() {
	_, playground := os.LookupEnv("GRAPHQL_PLAYGROUND")

	port, envPort := os.LookupEnv("PORT")
	if !envPort {
		port = "4567"
	}

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Query,
	})

	handler := handler.New(&handler.Config{
		Schema:     &schema,
		Playground: playground,
	})

	fmt.Println("Starting COVID-19 GraphQL Service!...")
	fmt.Println("Server listening on http://localhost:$PORT/graphql")

	http.Handle("/graphql", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
