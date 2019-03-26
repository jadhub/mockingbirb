package controller

import (
	"context"
	"net/http"
	"strings"

	configDomain "go.aoe.com/mockingbirb/src/config/domain"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	// MockController ...
	MockController struct {
		logger         flamingo.Logger
		responder      *web.Responder
		configProvider configDomain.ConfigProvider
		mockConfig     string
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
	config *struct {
		MockConfig string `inject:"config:mockconfig"`
	},
) {
	c.logger = logger
	c.responder = responder
	c.configProvider = configProvider
	c.mockConfig = config.MockConfig
}

func (c *MockController) GetConfigAction(ctx context.Context, req *web.Request) web.Result {
	configTree := c.configProvider.GetConfigTree()

	res := MockResult{
		Config: configTree,
	}

	return c.responder.Data(res).Status(http.StatusOK)
}

func (c *MockController) MockAction(ctx context.Context, req *web.Request) web.Result {
	configTree := c.configProvider.GetConfigTree()

	requestUrlPath := req.Request().URL.Path
	requestMethod := req.Request().Method

	responseConfig := c.getResponseConfig(configTree, requestUrlPath, requestMethod)

	c.responder.Data(responseConfig.Body)

	responseHeader := http.Header{}
	for key, value := range responseConfig.Headers {
		responseHeader[key] = make([]string, 1)
		responseHeader[key] = append(responseHeader[key], value)
	}

	response := c.responder.Data(responseConfig.Body).Status(uint(responseConfig.StatusCode))
	response.Header = responseHeader

	return response
}

func (c *MockController) getResponseConfig(tree configDomain.ConfigTree, requestUrlPath string, requestMethod string) *configDomain.Response {
	for _, config := range tree {
		for _, response := range config.Responses {
			if response.URI == requestUrlPath && strings.ToLower(response.Method) == strings.ToLower(requestMethod) {
				return &response
			}
		}
	}

	return nil
}
