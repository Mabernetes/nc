package handlers

import (
	"net/http"
	"node/logic"
)

type Status interface {
	Server(w http.ResponseWriter, r *http.Request)
	Deployment(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	logic  *logic.Logic
	Status Status
}

func NewHandler(logic *logic.Logic) *Handler {
	return &Handler{
		logic:  logic,
		Status: NewStatusHandler(logic),
	}
}
