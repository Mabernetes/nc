package main

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"log"
	"node/db"
	"node/handlers"
	"node/logic"
	"node/server"
)

func main() {
	dbConn, _ := sql.Open("postgres", "postgres:postgres@localhost:5432/nc?sslmode=disabledb")
	var DB *db.DB = db.NewDB(bun.NewDB(dbConn, pgdialect.New()))
	var Logic *logic.Logic = logic.NewLogic(DB)
	var Handler *handlers.Handler = handlers.NewHandler(Logic)
	log.Fatal(server.RunServer(Handler))
}
