package main

import (
	"github.com/RamazanBolatkhan/gin-gorm-rest/config"
	"github.com/RamazanBolatkhan/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run("localhost:9090")
}
