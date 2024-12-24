package handlers

import (
	"github.com/Mabernetes/nc/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusHandler struct {
	logic *services.Services
}

func NewStatusHandler(logic *services.Services) *StatusHandler {
	return &StatusHandler{logic: logic}
}

// Server
// @Summary      Show server status
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  services.ServerStatusData
// @Failure      500  {object}  interface{}
// @Router       /status/server [get]
func (h *StatusHandler) Server(ctx *gin.Context) {
	data := h.logic.Status.Server()

	ctx.JSON(http.StatusOK, data)
}

// Runner
// @Summary      Show runner status
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object} 	map[string]services.DeploymentStatusData
// @Failure      500  {object}  interface{}
// @Router       /status/runner [get]
func (h *StatusHandler) Runner(ctx *gin.Context) {
	data, err := h.logic.Status.Runner()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, data)
}
