package types

import (
	"github.com/graphql-go/graphql"
)

// CountryReport is a type definition.
var CountryReport = graphql.NewObject(graphql.ObjectConfig{
	Name:        "CountryReport",
	Description: "A country report total type.",
	Fields: graphql.Fields{
		"country": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A country location name.",
		},
		"provinces": &graphql.Field{
			Type:        graphql.NewList(Province),
			Description: "A country provices data.",
		},
		"latitude": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "A country location latitude.",
		},
		"longitude": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "A country location longitude.",
		},
		"date": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A date of query execution.",
		},
	},
})
