/*

Mockingbirb API Mock Server

Usage

 	go run main.go

*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"flamingo.me/flamingo/v3/core/canonicalurl"
	"flamingo.me/flamingo/v3/framework/cmd"

	// "flamingo.me/redirects"
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/flamingo/v3/core/zap"
	"flamingo.me/flamingo/v3/framework"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/prefixrouter"
	"go.aoe.com/mockingbirb/src"
	configuration "go.aoe.com/mockingbirb/src/config"
	mockingBirb "go.aoe.com/mockingbirb/src/server"
)

func main() {
	var mockingbirb = new(src.MOCKINGBIRB)

	rootContext := config.NewArea(
		"root",
		[]dingo.Module{
			new(framework.InitModule),
			new(cmd.Module),
			mockingbirb,
			new(zap.Module),
			new(prefixrouter.Module),
			new(requestlogger.Module),
			new(canonicalurl.Module),
			new(mockingBirb.Module),
			new(configuration.Module),
		},
		config.NewArea("main", []dingo.Module{}),
	)
	config.Load(rootContext, "config")

	defaultrouter := mockingbirb.Injector.GetInstance((*http.ServeMux)(nil)).(*http.ServeMux)
	addDefaultRoutes(defaultrouter)

	// has a reference to the injected RootCmd, so we can use it here :)
	if err := cmd.Run(rootContext.Injector); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func addDefaultRoutes(defaultRouter *http.ServeMux) {
	defaultRouter.HandleFunc("/ping", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(200)
		rw.Write([]byte("pong"))
	})

	defaultRouter.Handle("/favicon.ico", http.NotFoundHandler())

	defaultRouter.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Location", "/mockingbirb/")
		rw.WriteHeader(http.StatusTemporaryRedirect)
	})
}
