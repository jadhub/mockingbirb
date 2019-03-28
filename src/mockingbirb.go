package src

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/core/cache"
	"flamingo.me/flamingo/v3/framework/web"
)

// MOCKINGBIRB main setup
type MOCKINGBIRB struct {
	RouterRegistry *web.RouterRegistry `inject:""`
	Injector       *dingo.Injector     `inject:""`
}

// Configure mockingbirb application
func (mockingbirb *MOCKINGBIRB) Configure(injector *dingo.Injector) {
	injector.Bind((*cache.Backend)(nil)).ToInstance(cache.NewInMemoryCache())
}
