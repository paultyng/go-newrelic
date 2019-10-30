# go-newrelic [![Build Status](https://travis-ci.org/paultyng/go-newrelic.png?branch=master)](https://travis-ci.org/paultyng/go-newrelic)

go-newrelic is a Go SDK for communicating with New Relic APIs.

## Installation

Use `go get` to install go-newrelic.

```bash
$ go get -u github.com/paultyng/go-newrelic
```

## Usage

```go
import "github.com/paultyng/go-newrelic/v4/api"	// with go modules enabled (GO111MODULE=on or outside GOPATH)
import "github.com/paultyng/go-newrelic/api" // with go modules disabled
```

Construct a new New Relic client, then use the various methods on the client to access different parts of the [New Relic V2 REST API](https://docs.newrelic.com/docs/apis/rest-api-v2). For example:

```go
config := newrelic.Config{
	APIKey: "<NEWRELIC_API_KEY>",
}

client := newrelic.New(config)
apps, err := client.ListApplications()
```

### Infrastructure Alert Conditions

Infrastructure Alert Conditions are backed by the [New Relic Infrastructure API](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/infrastructure-alert-conditions/rest-api-calls-new-relic-infrastructure-alerts), and require the use of the `InfraClient` type:

```go
config := newrelic.Config{
	APIKey: "<NEWRELIC_API_KEY>",
}

client := newrelic.NewInfraClient(config)
apps, err := client.ListAlertInfraConditions(policyId)
```

The GoDoc link below details the available client options and the full list of available client methods.

## Contributing
Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Add some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

### Running tests
```bash
go test -v ./...
```

## License
[Apache-2.0](https://github.com/paultyng/go-newrelic/blob/master/LICENSE)

[![GoDoc](https://godoc.org/github.com/paultyng/go-newrelic?status.svg)](https://godoc.org/github.com/paultyng/go-newrelic)
