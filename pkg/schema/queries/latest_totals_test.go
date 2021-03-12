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

var latestTotal = `[{"confirmed":118501646,"recovered":93070573,"critical":89625,"deaths":3262812,"lastChange":"2021-03-08T15:26:23+01:00","lastUpdate":"2021-03-08T15:30:03+01:00"}]`

func TestLatestTotalsWithInvalidEndpoint(t *testing.T) {
	server := testutil.NewTestServer("totals", latestTotal, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", "")

	_, err := queries.LatestTotals.Resolve(graphql.ResolveParams{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
}

func TestLatestTotalsWithInvalidStatusCode(t *testing.T) {
	server := testutil.NewTestServer("totals", `{"message": "You are not subscribed to this API."}`, 401)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	_, err := queries.LatestTotals.Resolve(graphql.ResolveParams{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "You are not subscribed to this API.")
}

func TestLatestTotalsResult(t *testing.T) {
	server := testutil.NewTestServer("totals", latestTotal, 200)
	defer server.Close()

	os.Setenv("COVID_SERVICE_ENDPOINT", server.URL)

	result, _ := queries.LatestTotals.Resolve(graphql.ResolveParams{})
	assert.NotNil(t, result)

	model := result.(models.Total)
	assert.Equal(t, model.Confirmed, 118501646)
	assert.Equal(t, model.Recovered, 93070573)
	assert.Equal(t, model.Critical, 89625)
	assert.Equal(t, model.Deaths, 3262812)
}
