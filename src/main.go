package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/zap"
	"flamingo.me/flamingo/v3/framework/config"
	mockingBirbConfig "mockingbirb/src/mockconfig"
	mockingBirbServer "mockingbirb/src/mockserver"
)

type (
	mockingBirb struct {
		apiPort int
	}
)

// Inject dependencies
func (m *mockingBirb) Inject(
	cfg *struct {
		APIPort float64 `inject:"config:api.port"`
	},
) {
	m.apiPort = int(cfg.APIPort)
}

// DefaultConfig for this module
func (m *mockingBirb) DefaultConfig() config.Map {
	return config.Map{
		"api": config.Map{
			"port": 8080,
		},
	}
}

// Configure DI
func (m *mockingBirb) Configure(injector *dingo.Injector) {
	injector.Bind((*cache.Backend)(nil)).ToInstance(cache.NewInMemoryCache())
}

func main() {
	flamingo.App(
		[]dingo.Module{
			new(zap.Module),
			new(mockingBirbConfig.Module),
			new(mockingBirbServer.Module),
			new(mockingBirb),
		},
	)
}
