package server

import (
	"flamingo.me/dingo"

	configDomain "mockingbirb/src/mockconfig/domain"
	"mockingbirb/src/mockconfig/infrastructure"
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
