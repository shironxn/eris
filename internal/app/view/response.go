package view

import "github.com/gin-gonic/gin"

func JSON(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}
