package controller

import (
	"context"
	"net/http"

	configDomain "go.aoe.com/mockingbirb/src/config/domain"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	// MainController ...
	MainController struct {
		logger         flamingo.Logger
		responder      *web.Responder
		configProvider configDomain.ConfigProvider
		mockConfig     string
	}
	// Result ...
	Result struct {
		Config configDomain.ConfigTree
	}
)

func (c *MainController) Inject(
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

func (c *MainController) GetConfigAction(ctx context.Context, req *web.Request) web.Result {
	configTree := c.configProvider.GetConfigTree()

	res := Result{
		Config: configTree,
	}

	return c.responder.Data(res).Status(http.StatusOK)
}
