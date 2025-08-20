package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pm-cloudify/multi-container-app/server/api"
	"github.com/pm-cloudify/multi-container-app/server/internal/db"
)

func main() {
	// TODO: load these configs from env
	// init db connection

	// TODO: these are test user pass - it should be a secret
	mongodb_addr := "localhost:27017"
	port := "8080"
	mongodb_user := "root"
	mongodb_pass := "password"

	if val := os.Getenv("MONGODB_ADDR"); val != "" {
		mongodb_addr = val
	}
	if val := os.Getenv("GIN_PORT"); val != "" {
		port = val
	}
	if val := os.Getenv("MONGODB_USER"); val != "" {
		mongodb_user = val
	}
	if val := os.Getenv("MONGODB_PASS"); val != "" {
		mongodb_pass = val
	}

	db.ConnectToDB(mongodb_user + ":" + mongodb_pass + "@" + mongodb_addr)
	defer db.DisconnectDB()
	db.PingDB()

	// config routes
	router := gin.Default()
	api.AttachHandler(router)

	// launch server
	router.Run(":" + port)
}
