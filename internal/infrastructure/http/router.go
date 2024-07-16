package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/view"
)

type Router struct{}

func (r *Router) Route() http.Handler {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		view.JSON(ctx, http.StatusOK, gin.H{"message": "pong"})
	})

	return router
}
