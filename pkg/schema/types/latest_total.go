package types

import (
	"github.com/graphql-go/graphql"
)

// LatestTotal is a type definition.
var LatestTotal = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LatestTotal",
	Description: "A world latest total type.",
	Fields: graphql.Fields{
		"confirmed": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A total of confirmed cases.",
		},
		"recovered": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A total of recovered cases.",
		},
		"critical": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A total of critical cases.",
		},
		"deaths": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "A total of deaths.",
		},
		"lastChange": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A last change datetime.",
		},
		"lastUpdate": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A last update datetime.",
		},
	},
})
