package queries

import (
	"covid-19-graphql-service/pkg/models"
	"covid-19-graphql-service/pkg/schema/types"
	"covid-19-graphql-service/pkg/services"
	"covid-19-graphql-service/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
)

// SouthAmericaCountries a list of south america countries.
var SouthAmericaCountries = []string{
	"Argentina",
	"Bolivia",
	"Brazil",
	"Chile",
	"Colombia",
	"Ecuador",
	"Falkland Islands",
	"French Guiana",
	"Guyana",
	"Paraguay",
	"Peru",
	"Suriname",
	"Uruguay",
	"Venezuela",
}

// CountryReport query to retrieve a world latest totals.
var CountryReport = graphql.Field{
	Type:        types.CountryReport,
	Description: "A country report totals query.",
	Args: graphql.FieldConfigArgument{
		"country": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "An argument for country name.",
		},
		"date": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "An argument for report date.",
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		country := params.Args["country"].(string)
		date := params.Args["date"].(string)

		hostname, _ := os.LookupEnv("COVID_SERVICE_ENDPOINT")
		endpoint := fmt.Sprintf("%s/%s?name=%s&date=%s", hostname, "report/country/name", country, date)

		if !utils.SliceContains(SouthAmericaCountries, country) {
			return nil, errors.New("the country arg contains an invalid value")
		}

		response, err := services.GetCovidData(endpoint)
		if err != nil {
			return nil, err
		}

		if response.StatusCode != http.StatusOK {
			m := struct{ Message string }{}
			json.NewDecoder(response.Body).Decode(&m)

			return nil, errors.New(m.Message)
		}

		data := make([]models.Location, 1)
		json.NewDecoder(response.Body).Decode(&data)

		return data[0], nil
	},
}
