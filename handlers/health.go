package handlers

import (
	"github.com/PedroDrago/DogsAPI/types"
	"github.com/gin-gonic/gin"
)

func (h *HandlerController) HealthGetHandler(ctx *gin.Context) {
	ctx.JSON(200, types.JSON{
		"message": "Ok2",
	})
}
