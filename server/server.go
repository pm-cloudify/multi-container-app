package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pm-cloudify/multi-container-app/server/api"
	"github.com/pm-cloudify/multi-container-app/server/internal/db"
)

func main() {
	// TODO: load these configs from env
	// init db connection
	db.ConnectToDB("root:password@localhost:27017")
	defer db.DisconnectDB()
	db.PingDB()

	// config routes
	router := gin.Default()
	api.AttachHandler(router)

	// launch server
	router.Run(":8080")
}
