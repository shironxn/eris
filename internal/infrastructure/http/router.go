package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/view"
)

type router struct {
}

func (r *router) Route() http.Handler {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		view.JSON(ctx, http.StatusOK, gin.H{"message": "pong"})
	})

	// user Router

	v1 := router.Group("/api/v1")

	user := v1.Group("/user")
	user.POST("/")

	return router
}
