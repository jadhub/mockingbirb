# Mockingbirb

Mockingbirb is an HTTP API Mock server for defining http routes with predefined responses.

## Getting Started

- Install Go
- fire up ```CONTEXT=main go run mockingbirb.go serve```
- Its running!

### Prerequisites

You will need Golang in version 1.12.1 installed.

### Configuration

Currently, the project has a simple JSON config provider implemented that reads a json file and adds the contained routes to its internal config.

For an example, you should have a look at src/server/spec/test_json_config/config.json (which is also used for the test suite) as it
shows a few usecases how requests and responses should be structured. "matcherconfig" filters the requests to mockingbirb and the attached 
"responseconfig" sets the format of the response.

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Use the following to run the tests:

```
go test ./... -v
```

## Built With

* [Flamingo](https://go.aoe.com/#Home) - Scalable frontend framework for your headless microservice architecture

## Authors

* **Joachim Adomeit** - *Initial work* - [jadhub](https://github.com/jadhub)

## License

This project is licensed under the Open Software License v. 3.0 (OSL-3.0) - see the [LICENSE](LICENSE) file for details
