package server

import (
	goji "goji.io"
	"goji.io/pat"
	"node/handlers"
)

func StatusRouter(handler *handlers.Handler) *goji.Mux {
	subMux := goji.SubMux()

	subMux.HandleFunc(pat.Get("/server"), handler.Status.Server)
	subMux.HandleFunc(pat.Get("/runner"), handler.Status.Deployment)

	return subMux
}
