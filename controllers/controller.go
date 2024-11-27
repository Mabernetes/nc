package controllers

import (
	"github.com/gin-gonic/gin"
	"node/logic"
)

type Status interface {
	Server(ctx *gin.Context)
	Runner(ctx *gin.Context)
}

type Config interface {
	GetTree(ctx *gin.Context)
	GetConfig(ctx *gin.Context)
	UpdateConfig(ctx *gin.Context)
}

type Controller struct {
	logic  *logic.Logic
	Status Status
	Config Config
}

func NewController(logic *logic.Logic) *Controller {
	return &Controller{
		logic:  logic,
		Status: NewStatusController(logic),
		Config: NewConfigController(logic),
	}
}
