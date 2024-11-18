package server

import (
	"goji.io"
	"goji.io/pat"
	"net/http"
	"node/handlers"
	middlewares "node/middleware"
)

func RunServer(handler *handlers.Handler) error {
	mux := goji.NewMux()
	mux.Use(middlewares.LoggerMiddleware)
	mux.Use(middlewares.JSONResponseMiddleware)

	mux.Handle(pat.New("/status/*"), StatusRouter(handler))

	return http.ListenAndServe(":80", mux)
}
