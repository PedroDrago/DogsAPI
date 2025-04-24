package utils

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	return router
}
