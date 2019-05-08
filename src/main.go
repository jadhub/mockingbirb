package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/core/zap"
	"flamingo.me/flamingo/v3/framework"
	"flamingo.me/flamingo/v3/framework/cmd"
	mockingBirbConfig "mockingbirb/src/mockconfig"
	mockingBirbServer "mockingbirb/src/mockserver"
)

type (
	mockingBirb struct {}
)

// Inject dependencies
func (m *mockingBirb) Inject() {}

// Configure DI
func (m *mockingBirb) Configure(injector *dingo.Injector) {
	injector.Bind((*cache.Backend)(nil)).ToInstance(cache.NewInMemoryCache())
}

func main() {
	flamingo.App(
		[]dingo.Module{
			new(framework.InitModule),
			new(cmd.Module),
			new(zap.Module),
			new(mockingBirbConfig.Module),
			new(mockingBirbServer.Module),
			new(mockingBirb),
		},
	)
}
