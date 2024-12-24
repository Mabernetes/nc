package handlers

import (
	"github.com/Mabernetes/nc/src/services"
	"github.com/gin-gonic/gin"
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
	logic  *services.Services
	Status Status
	Config Config
}

func NewController(logic *services.Services) *Controller {
	return &Controller{
		logic:  logic,
		Status: NewStatusHandler(logic),
		Config: NewConfigHandler(logic),
	}
}
