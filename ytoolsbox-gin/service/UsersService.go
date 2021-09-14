package service

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {

	ctx.JSON(200, "hi")
}
