package types

import (
	"github.com/graphql-go/graphql"
)

// Province is a type definition.
var Province = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Province",
	Description: "A country location province.",
	Fields: graphql.Fields{
		"province": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A country province name.",
		},
		"confirmed": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A province total of confirmed cases.",
		},
		"recovered": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A province total of recovered cases.",
		},
		"deaths": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A province total of deaths.",
		},
		"active": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A province total of active cases.",
		},
	},
})
