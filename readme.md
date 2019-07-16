[![Build Status](https://api.travis-ci.org/jadhub/mockingbirb.svg?branch=master)](https://travis-ci.org/jadhub/mockingbirb)

# Mockingbirb

Mockingbirb is an HTTP API Mock server for defining http routes with predefined responses.

## Getting Started

- Install Go
- copy ```src/mockserver/spec/test_json_config/config.json``` to the ```mock_config``` folder and adapt it to your needed configuration
- fire up ```CONTEXT=dev make serve```
- Its running!

### Prerequisites

You will need Golang in version 1.12.1 installed.

### Configuration

Currently, the project has a simple JSON config provider implemented that reads a json file and adds the contained routes to its internal config.

For an example, you should have a look at src/mockserver/spec/test_json_config/config.json (which is also used for the test suite) as it shows a few usecases how requests and responses should be structured. "matcherconfig" filters the requests to mockingbirb and the attached 
"responseconfig" sets the format of the response.

Basically, a config looks like this:
`
[
  {
    "responses": [
      {
        "matcherconfig": {
          "uri": "/greet/me",
          "method": "get"
        },
        "responseconfig": {
          "statusCode": 200,
          "headers": {
            "Content-Type": "text/plain; charset=utf-8"
          },
          "body": {
            "message": "Mockingbirb says hi!"
          }
        }
      }
		]
  }
]
`
A matcherconfig instructs mockingbirb to react to a request to /greet/me, but only if its a GET Request. The responseconfig shapes the response to the request, answering with HTTP Status 200, a text/plain Content-Type and a nice Message in the Response Body. There are more examples in the test config, read up for further ideas, i.e. specific parameters in the matcherconfig. 

## Running the tests

Use the following to run the tests:

```
make test
```

## Built With

* [Flamingo](https://go.aoe.com/#Home) - Scalable frontend framework for your headless microservice architecture

## Authors

* **Joachim Adomeit** - *Initial work* - [jadhub](https://github.com/jadhub)

## License

This project is licensed under the Open Software License v. 3.0 (OSL-3.0) - see the [LICENSE](LICENSE) file for details

