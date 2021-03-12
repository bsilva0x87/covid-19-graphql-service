package queries_test

import (
	"covid-19-graphql-service/pkg/models"
	"covid-19-graphql-service/pkg/schema/queries"
	"covid-19-graphql-service/test/testutil"
	"os"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("COVID_SERVICE_ENDPOINT", "http://localhost")
	os.Setenv("COVID_SERVICE_API_KEY", "fake-api-key-0xf56")
}

var latestCountryData = `[{"country": "Brazil", "code": "BR", "confirmed": 11125017, "recovered": 9843218, "critical": 8318, "deaths": 268568, "latitude": -14.235004, "longitude": -51.92528, "lastChange": "2021-03-10T00:29:45+01:00", "lastUpdate": "2021-03-10T19:00:04+01:00"}]`

func TestLatestCountryDataWithInvalidEndpoint(t *testing.T) {
	server := testutil.NewTestServer("country", latestCountryData, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", "")

	_, err := queries.LatestCountryData.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"name": "Brazil"},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
}

func TestLatestCountryDataWithInvalidStatusCode(t *testing.T) {
	server := testutil.NewTestServer("country", `{"message": "You are not subscribed to this API."}`, 401)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	_, err := queries.LatestCountryData.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"name": "Brazil"},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "You are not subscribed to this API.")
}

func TestLatestCountryDataWithInvalidCountry(t *testing.T) {
	server := testutil.NewTestServer("country", countryReport, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	_, err := queries.LatestCountryData.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"name": "XXX"},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "the country arg contains an invalid value")
}

func TestLatestCountryDataResult(t *testing.T) {
	server := testutil.NewTestServer("country", latestCountryData, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	result, _ := queries.LatestCountryData.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"name": "Brazil"},
	})

	assert.NotNil(t, result)

	model := result.(models.Country)
	assert.Equal(t, model.Confirmed, 11125017)
	assert.Equal(t, model.Recovered, 9843218)
	assert.Equal(t, model.Critical, 8318)
	assert.Equal(t, model.Deaths, 268568)
	assert.Equal(t, model.Latitude, -14.235004)
	assert.Equal(t, model.Longitude, -51.92528)
	assert.Equal(t, model.LastChange, "2021-03-10T00:29:45+01:00")
	assert.Equal(t, model.LastUpdate, "2021-03-10T19:00:04+01:00")
}
