package main

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "node/docs"
	"node/src/cron"
	"node/src/handlers"
	"node/src/logic"
	"time"
)

// @title           M8s Node Controller API
// @version         0.1

// @contact.name    Marsh Meg
// @contact.url    	https://t.me/marshmeg
// @contact.email  	uraevdmitrij031@gmail.com

// @license.name 	MI

// @host      		localhost:8000
// @BasePath  		/
func main() {
	// Инициализация слоёв
	var Logic *logic.Logic = logic.NewLogic()
	var Handler *handlers.Controller = handlers.NewController(Logic)

	// запуск Goroutine
	var tasks *cron.Cron = cron.NewCronHandler(Logic.Config)
	go cronRun(tasks, context.Background())

	// запуск API
	r := gin.Default()

	api := r.Group("/api")
	{
		statusR := api.Group("/status")
		{
			statusR.GET("server", Handler.Status.Server)
			statusR.GET("runner", Handler.Status.Runner)
		}

		configR := api.Group("/configs")
		{
			configR.GET("", Handler.Config.GetTree)
			configR.GET(":deployment/:pod", Handler.Config.GetConfig)
			configR.POST(":deployment/:pod", Handler.Config.UpdateConfig)
			configR.PUT(":deployment/:pod", Handler.Config.UpdateConfig)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":9000")
	if err != nil {
		panic(err)
	}
}

func cronRun(cron *cron.Cron, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			cron.Start()
		}

		time.Sleep(time.Minute)
	}
}
