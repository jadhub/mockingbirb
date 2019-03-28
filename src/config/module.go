package server

import (
	"flamingo.me/dingo"
	configDomain "go.aoe.com/mockingbirb/src/config/domain"
	"go.aoe.com/mockingbirb/src/config/infrastructure"
)

type (
	// Module ...
	Module struct{}
)

// Inject dependencies
func (m *Module) Inject() {}

// Configure module
func (m *Module) Configure(injector *dingo.Injector) {
	injector.Bind((*configDomain.ConfigProvider)(nil)).ToProvider(infrastructure.NewJSONConfigProvider).AsEagerSingleton()
}
