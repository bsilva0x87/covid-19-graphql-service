package services

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

// HTTPClient interface for covid data service.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client constant for http client.
var Client HTTPClient

func init() {
	Client = &http.Client{}
}

// GetCovidData function to retrieve covid world data.
func GetCovidData(endpoint string) (*http.Response, error) {
	log.Println("GET", endpoint)

	request, _ := http.NewRequest(http.MethodGet, endpoint, nil)

	url, _ := url.Parse(endpoint)
	key, _ := os.LookupEnv("COVID_SERVICE_API_KEY")

	request.Header.Add("x-rapidapi-key", key)
	request.Header.Add("x-rapidapi-host", url.Host)

	return Client.Do(request)
}
