package schema

import (
	"covid-19-graphql-service/pkg/schema/queries"

	"github.com/graphql-go/graphql"
)

// Query base definition.
var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestTotals":      &queries.LatestTotals,
		"latestCountryData": &queries.LatestCountryData,
		"countryReport":     &queries.CountryReport,
	},
})
