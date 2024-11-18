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
