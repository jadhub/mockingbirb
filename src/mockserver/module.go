package mockserver

import (
	"fmt"
	"strings"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"

	configDomain "mockingbirb/src/mockconfig/domain"
	"mockingbirb/src/mockserver/interfaces/controller"
)

type (
	// Module with routerRegistry
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

// Routes for mockingbirb api
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("mockingbirb.api.getConfig", r.mockController.GetConfigAction)
	_, err := registry.Route("/api/getConfig", "mockingbirb.api.getConfig")
	if err != nil {
		panic(fmt.Sprintf("unexpected route bind error: %v", err))
	}

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
			key := strings.Replace(response.MatcherConfig.URI, "/", ".", -1)

			registry.HandleAny(key, r.mockController.MockAction)
			_, err := registry.Route(response.MatcherConfig.URI, key)
			if err != nil {
				panic(fmt.Sprintf("unexpected route bind error: %v", err))
			}
		}
	}
}
