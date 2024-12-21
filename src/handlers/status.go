package handlers

import (
	"github.com/Mabernetes/nc/src/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusHandler struct {
	logic *logic.Logic
}

func NewStatusHandler(logic *logic.Logic) *StatusHandler {
	return &StatusHandler{logic: logic}
}

// Server
// @Summary      Show server status
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  logic.ServerStatusData
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
// @Success      200  {object} 	map[string]logic.DeploymentStatusData
// @Failure      500  {object}  interface{}
// @Router       /status/runner [get]
func (h *StatusHandler) Runner(ctx *gin.Context) {
	data, err := h.logic.Status.Runner()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, data)
}
