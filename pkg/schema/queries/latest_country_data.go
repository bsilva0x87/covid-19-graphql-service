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

// LatestCountryData query to retrieve a country latest totals.
var LatestCountryData = graphql.Field{
	Type: types.LatestCountryData,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "An argument for country name.",
		},
	},
	Description: "A latest country total cases data.",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name := params.Args["name"].(string)
		hostname, _ := os.LookupEnv("COVID_SERVICE_ENDPOINT")
		endpoint := fmt.Sprintf("%s/%s?name=%s", hostname, "country", name)

		response, err := services.GetCovidData(endpoint)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		if !utils.SliceContains(SouthAmericaCountries, name) {
			return nil, errors.New("the country arg contains an invalid value")
		}

		if response.StatusCode != http.StatusOK {
			var m ErrorMessage
			json.NewDecoder(response.Body).Decode(&m)

			return nil, errors.New(m.Messsage)
		}

		data := make([]models.Country, 1)
		json.NewDecoder(response.Body).Decode(&data)

		return data[0], nil
	},
}
