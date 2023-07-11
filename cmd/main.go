package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"repeatTestProject/config"
	"repeatTestProject/internal/db"
	"repeatTestProject/internal/repository"
	"repeatTestProject/internal/server"
	"repeatTestProject/internal/service"
)

func main() {
	config.ReadConfig("./config/config.json")

	//DB connection with Repository
	dbConn := db.ConnectionToDB()
	repo := repository.NewRepo(dbConn)

	//Connection Repository with service
	serviceCon := service.NewService(repo)

	r := gin.Default()

	//Connection service with handle
	handle := server.NewHandler(serviceCon, r)

	handle.AllRoutes()

	err := r.Run(config.Configure.PortRun)
	if err != nil {
		log.Fatal("router failed to start")
	}

}
