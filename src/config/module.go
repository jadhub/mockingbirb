package server

import (
	"flamingo.me/dingo"
	configApplication "go.aoe.com/mockingbirb/src/config/application"
	configDomain "go.aoe.com/mockingbirb/src/config/domain"
)

type (
	// Module ...
	Module struct{}
)

// Inject dependencies
func (m *Module) Inject() {}

// Configure module
func (m *Module) Configure(injector *dingo.Injector) {
	injector.Bind((*configDomain.ConfigProvider)(nil)).To(new(configApplication.JsonConfigProvider))
}
