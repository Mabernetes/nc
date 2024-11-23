package server

import (
	goji "goji.io"
	"goji.io/pat"
	"node/handlers"
)

func ConfigRouter(handler *handlers.Handler) *goji.Mux {
	subMux := goji.SubMux()

	subMux.HandleFunc(pat.Get("/"), handler.Config.ServeHTTP)

	return subMux
}
