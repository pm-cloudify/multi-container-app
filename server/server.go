package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pm-cloudify/multi-container-app/server/api"
)

func main() {
	router := gin.Default()
	api.AttachHandler(router)
	router.Run(":8080")
}
