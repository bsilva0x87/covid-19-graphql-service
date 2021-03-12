package queries

import (
	"covid-19-graphql-service/pkg/models"
	"covid-19-graphql-service/pkg/schema/types"
	"covid-19-graphql-service/pkg/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
)

// ErrorMessage struct to represent response body error message.
type ErrorMessage struct {
	Messsage string `json:"message"`
}

// LatestTotals query to retrieve a world latest totals.
var LatestTotals = graphql.Field{
	Type:        types.LatestTotal,
	Description: "A world latest total query.",
	Resolve: func(_ graphql.ResolveParams) (interface{}, error) {
		hostname, _ := os.LookupEnv("COVID_SERVICE_ENDPOINT")
		endpoint := fmt.Sprintf("%s/%s", hostname, "totals")

		response, err := services.GetCovidData(endpoint)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			var m ErrorMessage
			json.NewDecoder(response.Body).Decode(&m)

			return nil, errors.New(m.Messsage)
		}

		data := make([]models.Total, 1)
		json.NewDecoder(response.Body).Decode(&data)

		return data[0], nil
	},
}
