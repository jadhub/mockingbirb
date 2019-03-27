package controller

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	configDomain "go.aoe.com/mockingbirb/src/config/domain"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	// MockController ...
	MockController struct {
		Logger         flamingo.Logger
		Responder      *web.Responder
		ConfigProvider configDomain.ConfigProvider
	}
	// MockResult ...
	MockResult struct {
		Config configDomain.ConfigTree
	}
)

func (c *MockController) Inject(
	responder *web.Responder,
	logger flamingo.Logger,
	configProvider configDomain.ConfigProvider,
) {
	c.Logger = logger
	c.Responder = responder
	c.ConfigProvider = configProvider
}

func (c *MockController) GetConfigAction(ctx context.Context, req *web.Request) web.Result {
	configTree := c.ConfigProvider.GetConfigTree()

	res := MockResult{
		Config: configTree,
	}

	return c.Responder.Data(res).Status(http.StatusOK)
}

func (c *MockController) MockAction(ctx context.Context, req *web.Request) web.Result {
	configTree := c.ConfigProvider.GetConfigTree()

	responseConfig := c.getResponseConfig(configTree, req)

	c.Responder.Data(responseConfig.ResponseConfig.Body)

	responseHeader := http.Header{}

	for key, value := range responseConfig.ResponseConfig.Headers {
		responseHeader[key] = make([]string, 1)
		responseHeader[key] = append(responseHeader[key], value)
	}

	responseBody := ""
	val, ok := responseConfig.ResponseConfig.Body.(string)
	if !ok {
		response := c.Responder.Data(responseConfig.ResponseConfig.Body).Status(uint(responseConfig.ResponseConfig.StatusCode))
		response.Header = responseHeader

		return response
	}

	responseBody = val

	return &web.Response{
		Status: uint(responseConfig.ResponseConfig.StatusCode),
		Header: responseHeader,
		Body:   bytes.NewBufferString(responseBody),
	}
}

func (c *MockController) getResponseConfig(tree configDomain.ConfigTree, req *web.Request) *configDomain.Response {
	// Request Data
	requestURI := req.Request().URL.Path
	requestMethod := req.Request().Method
	requestGetParams := req.QueryAll()
	requestPostParams, err := req.FormAll()
	if err != nil {
		requestPostParams = nil
	}

	for _, config := range tree {
		for _, response := range config.Responses {
			// matcher Data
			matcherURI := response.MatcherConfig.URI
			matcherMethod := response.MatcherConfig.Method
			matcherGetParams := response.MatcherConfig.Params.GET
			matcherPostParams := response.MatcherConfig.Params.POST

			// check if URI and Http method match
			if matcherURI == requestURI && strings.ToLower(matcherMethod) == strings.ToLower(requestMethod) {
				// check if matcherGetParams are set
				fmt.Printf("%d", len(requestGetParams))
				if len(matcherGetParams) > 0 {
					getValidation := true

					for key, value := range matcherGetParams {
						// check if matcherGetParams don't match request
						if requestGetParams.Get(key) != value {
							getValidation = false
						}
					}

					if getValidation == true {
						return &response
					} else {
						return c.ParamMismatchResponse()
					}
				}

				// check if matcherPostParams are set
				if len(matcherPostParams) > 0 {
					postValidation := false

					for key, matcherValue := range matcherPostParams {
						// Check if matcherGetParams don't match request
						for _, paramValue := range requestPostParams[key] {
							if paramValue == matcherValue {
								postValidation = true
							}
						}
					}

					if postValidation == true {
						return &response
					} else {
						return c.ParamMismatchResponse()
					}
				}

				// matcherGetParams and matcherPostParams are optional, if URI and http method match, return response
				return &response
			}
		}
	}

	return c.NoConfigFoundResponse()
}

func (c *MockController) ParamMismatchResponse() *configDomain.Response {
	return &configDomain.Response{
		ResponseConfig: struct {
			StatusCode int               `json:"statusCode"`
			Headers    map[string]string `json:"headers"`
			Body       interface{}       `json:"body,omitempty"`
		}{
			StatusCode: 404,
			Headers: map[string]string{
				"Content-Type": "text/plain; charset=utf-8",
			},
			Body: "Mockingbirb config found, but param mismatch",
		},
	}
}

func (c *MockController) NoConfigFoundResponse() *configDomain.Response {
	return &configDomain.Response{
		ResponseConfig: struct {
			StatusCode int               `json:"statusCode"`
			Headers    map[string]string `json:"headers"`
			Body       interface{}       `json:"body,omitempty"`
		}{
			StatusCode: 404,
			Headers: map[string]string{
				"Content-Type": "text/plain; charset=utf-8",
			},
			Body: "Mockingbirb config not found for this request",
		},
	}
}
