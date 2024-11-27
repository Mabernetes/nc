package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"node/logic"
)

type StatusController struct {
	logic *logic.Logic
}

func NewStatusController(logic *logic.Logic) *StatusController {
	return &StatusController{logic: logic}
}

// Server
// @Summary      Show server status
// @Tags         status
// @Accept       json
// @Produce      json
// @Success      200  {object}  logic.ServerStatusData
// @Failure      500  {object}  interface{}
// @Router       /status/server [get]
func (h *StatusController) Server(ctx *gin.Context) {
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
func (h *StatusController) Runner(ctx *gin.Context) {
	data, err := h.logic.Status.Runner()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, data)
}
