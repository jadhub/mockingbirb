package server

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
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
	controller *controller.MainController
}

// Routes for cart api
func (r *routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("mockingbirb.api.getConfig", r.controller.GetConfigAction)
	registry.Route("/api/getConfig", "mockingbirb.api.getConfig")
}

// Inject method
func (r *routes) Inject(
	apiController *controller.MainController,
) {
	r.controller = apiController
}
