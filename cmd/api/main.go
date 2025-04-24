package main

import (
	// "fmt"
	// "net/http"

	// "github.com/PedroDrago/DogsAPI/data"
	"github.com/PedroDrago/DogsAPI/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	controller := handlers.HandlerController{}alskdjf
	router := gin.Default()
	router.SetTrustedProxies(nil)
	routerV1 := router.Group("/v1")
	routerV1.GET("health", controller.HealthGetHandler)
	router.Run()
}
