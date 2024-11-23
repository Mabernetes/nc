package handlers

import (
	"encoding/json"
	"net/http"
	"node/logic"
)

type ConfigHandler struct {
	logic *logic.Logic
}

func NewConfigHandler(logic *logic.Logic) *ConfigHandler {
	return &ConfigHandler{logic: logic}
}

func (h *ConfigHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := h.logic.Config.GetTree()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
