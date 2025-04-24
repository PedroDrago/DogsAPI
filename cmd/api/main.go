package main

import (
	"os"

	"github.com/PedroDrago/DogsAPI/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	controller := handlers.HandlerController{}
	router := gin.Default()
	routerV1 := router.Group("/v1")
	routerV1.GET("health", controller.HealthGetHandler)
	router.Run()
}
