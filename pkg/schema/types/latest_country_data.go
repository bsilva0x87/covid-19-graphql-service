package types

import (
	"github.com/graphql-go/graphql"
)

// LatestCountryData is a type definition.
var LatestCountryData = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LatestCountryData",
	Description: "A latest country data type.",
	Fields: graphql.Fields{
		"country": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A country name.",
		},
		"code": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "A country code.",
		},
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
		"latitude": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "A country location latitude.",
		},
		"longitude": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "A country location longitude.",
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
