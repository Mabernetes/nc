package main

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"node/controllers"
	"node/cron"
	"node/db"
	_ "node/docs"
	"node/logic"
	"time"
)

// @title           M8s Node Controller API
// @version         0.1
// @description     This is a sample server seller server.
// @termsOfService  http://swagger.io/terms/
// @contact.name    Marsh Meg
// @contact.url    	https://t.me/marshmeg
// @contact.email  	uraevdmitrij031@gmail.com
// @license.name 	MI
// @host      		localhost:8000
// @BasePath  		/
func main() {
	// инициализация слоёв
	dbConn, _ := sql.Open("postgres", "postgres:postgres@localhost:5432/nc?sslmode=disabledb")
	var DB *db.DB = db.NewDB(bun.NewDB(dbConn, pgdialect.New()))
	var Logic *logic.Logic = logic.NewLogic(DB)
	var Handler *controllers.Controller = controllers.NewController(Logic)

	// запуск Goroutine
	var tasks *cron.Cron = cron.NewCronHandler(Logic.Config)
	go cronRun(tasks, context.Background())

	// запуск API
	r := gin.Default()
	{
		statusR := r.Group("/status")
		{
			statusR.GET("server", Handler.Status.Server)
			statusR.GET("runner", Handler.Status.Runner)
		}

		configR := r.Group("/configs")
		{
			configR.GET("", Handler.Config.GetTree)
			configR.GET(":deployment/:pod", Handler.Config.GetConfig)
			configR.POST(":deployment/:pod", Handler.Config.UpdateConfig)
			configR.PUT(":deployment/:pod", Handler.Config.UpdateConfig)
		}

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	err := r.Run(":8000")
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
