package services_test

import (
	"covid-19-graphql-service/pkg/services"
	"covid-19-graphql-service/test/testutil"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

var json = `[{"confirmed":118501646,"recovered":93070573,"critical":89625,"deaths":3262812,"lastChange":"2021-03-08T15:26:23+01:00","lastUpdate":"2021-03-08T15:30:03+01:00"}]`

func TestGetCovidData(t *testing.T) {
	server := testutil.NewTestServer("totals", json, 400)
	defer server.Close()

	response, err := services.GetCovidData(fmt.Sprintf("%s/%s", server.URL, "totals"))
	assert.NoError(t, err)

	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, string(body), json)
}
