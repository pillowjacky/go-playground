package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"gitlab.com/cwkj/cloud-native/go-application/route"
)

func InjectRouter() *chi.Mux {
	panic(wire.Build(
		route.ProvideRouter,
	))
}
