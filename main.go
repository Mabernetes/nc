package main

import (
	_ "github.com/Mabernetes/nc/docs"
	"github.com/Mabernetes/nc/src/cron"
	"github.com/Mabernetes/nc/src/handlers"
	"github.com/Mabernetes/nc/src/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           M8s Node Controller API
// @version         0.1

// @contact.name    Marsh Meg
// @contact.url    	https://t.me/marshmeg
// @contact.email  	uraevdmitrij031@gmail.com

// @license.name 	MI

// @host      		localhost:9000
// @BasePath  		/api
func main() {
	// Инициализация слоёв
	var layerServices *services.Services = services.NewLogic()
	var layerHandlers *handlers.Controller = handlers.NewController(layerServices)

	// запуск Goroutine
	go func() {
		cron.Start(layerServices)
	}()

	// запуск API
	r := gin.Default()

	api := r.Group("/api")
	{
		statusR := api.Group("/status")
		{
			statusR.GET("server", layerHandlers.Status.Server)
			statusR.GET("runner", layerHandlers.Status.Runner)
		}

		configR := api.Group("/configs")
		{
			configR.GET("", layerHandlers.Config.GetTree)
			configR.GET(":deployment/:pod", layerHandlers.Config.GetConfig)
			configR.POST(":deployment/:pod", layerHandlers.Config.UpdateConfig)
			configR.PUT(":deployment/:pod", layerHandlers.Config.UpdateConfig)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":9000")
	if err != nil {
		panic(err)
	}
}
