package handlers

import (
	"encoding/json"
	"net/http"
	"node/logic"
)

type StatusHandler struct {
	logic *logic.Logic
}

func NewStatusHandler(logic *logic.Logic) *StatusHandler {
	return &StatusHandler{logic: logic}
}

func (h *StatusHandler) Server(w http.ResponseWriter, r *http.Request) {
	load := h.logic.Status.Server()

	data, err := json.Marshal(load)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (h *StatusHandler) Deployment(w http.ResponseWriter, r *http.Request) {
	var deployment, pod string
	deployment = r.URL.Query().Get("deployment")
	pod = r.URL.Query().Get("pod")

	data, err := h.logic.Status.Deployment(deployment, pod)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	byteData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(byteData)
}
