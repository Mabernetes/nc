package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	"node/logic"
	"node/utils"
)

type ConfigController struct {
	logic *logic.Logic
}

func NewConfigController(logic *logic.Logic) *ConfigController {
	return &ConfigController{logic: logic}
}

// GetTree
// @Summary      Show configs tree
// @Tags         config
// @Accept       json
// @Produce      json
// @Success      200  {object}  logic.ConfigsTree
// @Failure      500  {object}  interface{}
// @Router       /configs/ [get]
func (h *ConfigController) GetTree(ctx *gin.Context) {
	data, err := h.logic.Config.GetTree()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

// GetConfig
// @Summary      Show config file
// @Tags         config
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.ComposeFile
// @Failure      500  {object}  interface{}
// @Router       /configs/:deployment/:pod [get]
func (h *ConfigController) GetConfig(ctx *gin.Context) {
	var deployment, pod string
	deployment = ctx.Param("deployment")
	pod = ctx.Param("pod")

	str, _ := h.logic.Config.ReadConfigFile(deployment, pod)
	ctx.JSON(http.StatusOK, str)
}

// UpdateConfig
// @Summary      Save or update config file
// @Description	 POST == PUT
// @Tags         config
// @Accept       json
// @Produce      json
// @Success      200  {object}  nil
// @Failure      500  {object}  interface{}
// @Router       /configs/:deployment/:pod [post]
// @Router       /configs/:deployment/:pod [put]
func (h *ConfigController) UpdateConfig(ctx *gin.Context) {
	var deployment, pod string
	var data utils.ComposeFile
	deployment = ctx.Param("deployment")
	pod = ctx.Param("pod")
	err := ctx.ShouldBindBodyWithJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = h.logic.Config.SaveConfigFile(deployment, pod, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	str, _ := h.logic.Config.ReadConfigFile(deployment, pod)
	ctx.JSON(http.StatusOK, str)
}