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

var countryReport = `[{"country": "Brazil","date": "2020-01-01","latitude": -14.235004,"longitude": -51.92528,"provinces": [{"active": 0,"confirmed": 0,"deaths": 0,"province": "Brazil"}]}]`

func TestCountryReportWithInvalidEndpoint(t *testing.T) {
	server := testutil.NewTestServer("report/country/name", countryReport, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", "")

	_, err := queries.CountryReport.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"country": "Brazil", "date": "2020-01-01"},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
}

func TestCountryReportWithInvalidStatusCode(t *testing.T) {
	server := testutil.NewTestServer("report/country/name", `{"message": "You are not subscribed to this API."}`, 401)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	_, err := queries.CountryReport.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"country": "Brazil", "date": "2020-01-01"},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "You are not subscribed to this API.")
}

func TestCountryReportWithInvalidCountry(t *testing.T) {
	server := testutil.NewTestServer("report/country/name", countryReport, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	_, err := queries.CountryReport.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"country": "XXX", "date": "2020-01-01"},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "the country arg contains an invalid value")
}

func TestCountryReportResult(t *testing.T) {
	server := testutil.NewTestServer("report/country/name", countryReport, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	result, _ := queries.CountryReport.Resolve(graphql.ResolveParams{
		Args: map[string]interface{}{"country": "Brazil", "date": "2020-01-01"},
	})

	assert.NotNil(t, result)

	model := result.(models.Location)
	assert.Equal(t, model.Country, "Brazil")
	assert.Equal(t, model.Date, "2020-01-01")
	assert.Equal(t, model.Latitude, -14.235004)
	assert.Equal(t, model.Longitude, -51.92528)

	for _, province := range model.Provinces {
		assert.Equal(t, province.Active, 0)
		assert.Equal(t, province.Confirmed, 0)
		assert.Equal(t, province.Deaths, 0)
		assert.Equal(t, province.Recovered, 0)
		assert.Equal(t, province.Province, "Brazil")
	}
}
