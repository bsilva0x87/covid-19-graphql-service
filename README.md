# COVID-19 GraphQL Service
This is a basic implementation for the study proposal, of a COVID-19 report service using GraphQL and Golang.

## About project
The package `/cmd` covers a [main](cmd/) entrypoint:
* The basic form of an executable command and entry point.
* Importing packages (from the standard library and the remote repository)
* Using the GraphQL Go handler as request middleware ([graphql-go/handler](https://github.com/graphql-go/handler))
* Listen and serve command, to listen on the TCP network requests on :4567 port.

The package `/pkg` covers `models`, `schema`, `services` and `utils`:
* `pkg/models` covers struct type for data representation.
* `pkg/schema` covers GraphQL Schema definitions (Eg. Query, Types, etc).
* `pkg/services` covers external services to retrieve COVID-19 data from the Rapid API service ([COVID-19 data](https://rapidapi.com/Gramzivi/api/covid-19-data)).
* `pkg/utils` covers program helpers utility.

The package `/test` covers only `testutil`
* `test/testutil` covers an HTTP Server to mock request data, avoiding real work requests.

## How to execute
This project raise a new HTTP Server and listen to request on port :4567.


### Installing dependencies
This project was build using `go mod`. Then to install project dependencies, run:
```
go mod download
```

### Developer mode run
To execute this project in developer mode, clone this repository and just run:
```
go run cmd/main.go
```

### Production mode run
To execute this project as binary production mode, run:
```
go build cmd/main.go && ./main
```

### Test mode run
To execute this project as test mode, just go to a covered package (including `_test.go` files) and run:
```
go test -v
```
Optionally, you can use cover profile `go test ./pkg/ -coverprofile=coverage.out -v` and `go tool cover -html coverage.out` to open a package coverage browser version.

### Extra
This project uses `.env` to load environment variables. You'll need to create a new file on the root directory called `.env` before run. Fill this file with these values:
```
COVID_SERVICE_ENDPOINT=https://covid-19-data.p.rapidapi.com
COVID_SERVICE_API_KEY=get-your-rapid-api-key
GRAPHQL_PLAYGROUND=true
```
Optionally, you can use GraphQL Playground setting the enviroment variable `GRAPHQL_PLAYGROUND`as `true`.

&copy;bsilva0x87, 2021.