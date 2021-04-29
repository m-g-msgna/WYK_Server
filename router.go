package main

import (
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {
	sqlConn := DbConnectionInjector()

	wykController := ServiceContainer().InjectWYKController(sqlConn)
	

	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
	)

	// [Initialize new user] .../wyk/api/v1/user/initialize
	r.Post("/wyk/api/v1/user/initialize", wykController.Initialize )

	//[Update hash after contacts list is modifies] .../wyk/api/v1/user/update
	r.Put("/wyk/api/v1/user/update", wykController.Update)

	//[Authenticate a user via an application] .../wyk/api/v1/user/authenticate
	r.Post("/wyk/api/v1/user/authenticate", wykController.Authenticate)

	//[Get user data] .../wyk/api/v1/user/info?user-id={1000}
	r.Get("/wyk/api/v1/user/info", wykController.GetUserData)


	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
