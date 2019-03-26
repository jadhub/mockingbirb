package server

import (
	"strings"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	configDomain "go.aoe.com/mockingbirb/src/config/domain"
	"go.aoe.com/mockingbirb/src/server/interfaces/controller"
)

type (
	// Module ...
	Module struct {
		routerRegistry *web.RouterRegistry
	}
)

// Inject dependencies
func (m *Module) Inject(
	routerRegistry *web.RouterRegistry,
) {
	m.routerRegistry = routerRegistry
}

// Configure module
func (m *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(routes))
}

type routes struct {
	mockController *controller.MockController
	configProvider configDomain.ConfigProvider
}

// Routes for cart api
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("mockingbirb.api.getConfig", r.mockController.GetConfigAction)
	registry.Route("/api/getConfig", "mockingbirb.api.getConfig")

	r.registerMockRoutes(registry)
}

// Inject method
func (r *routes) Inject(
	mockController *controller.MockController,
	configProvider configDomain.ConfigProvider,
) {
	r.mockController = mockController
	r.configProvider = configProvider
}

func (r *routes) registerMockRoutes(registry *web.RouterRegistry) {
	for _, config := range r.configProvider.GetConfigTree() {
		for _, response := range config.Responses {
			key := strings.Replace(response.URI, "/", ".", -1)

			registry.HandleAny(key, r.mockController.MockAction)
			registry.Route(response.URI, key)
		}
	}
}
