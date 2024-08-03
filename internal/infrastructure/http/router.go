package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/controller"
	"github.com/shironxn/eris/internal/app/view"
)

type Router interface {
	Route() http.Handler
}

type router struct {
	user controller.UserController
}

func NewRouter(user controller.UserController) Router {
	return &router{
		user: user,
	}
}

func (r *router) Route() http.Handler {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		view.JSON(ctx, http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	user := v1.Group("/user")

	auth.POST("/login", r.user.Login)
	auth.POST("/register", r.user.Register)

	user.GET("/", r.user.GetAll)
	user.GET("/:id", r.user.GetByID)
	user.PUT("/:id", r.user.Update)
	user.DELETE("/:id", r.user.Delete)

	return router
}
