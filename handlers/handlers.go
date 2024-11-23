package handlers

import (
	"net/http"
	"node/logic"
)

type Status interface {
	Server(w http.ResponseWriter, r *http.Request)
	Deployment(w http.ResponseWriter, r *http.Request)
}

type Config interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	logic  *logic.Logic
	Status Status
	Config Config
}

func NewHandler(logic *logic.Logic) *Handler {
	return &Handler{
		logic:  logic,
		Status: NewStatusHandler(logic),
		Config: NewConfigHandler(logic),
	}
}
