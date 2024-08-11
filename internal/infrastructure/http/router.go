package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/controller"
)

type Router struct {
	Controller Controllers
	Middleware Middleware
}

type Controllers struct {
	User     controller.UserController
	Product  controller.ProductController
	Category controller.CategoryController
}

func NewRouter(controllers Controllers, middleware Middleware) *Router {
	return &Router{
		Controller: controllers,
		Middleware: middleware,
	}
}

func (r *Router) Route() http.Handler {
	router := gin.Default()

	router.Use(r.Middleware.Auth())

	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	user := v1.Group("/users")
	product := v1.Group("/products")
	category := v1.Group("/categories")

	// Auth routes
	auth.POST("/login", r.Controller.User.Login)
	auth.POST("/register", r.Controller.User.Register)

	// User routes
	user.GET("/", r.Controller.User.GetAll)
	user.GET("/:id", r.Controller.User.GetByID)
	user.PUT("/:id", r.Controller.User.Update)
	user.DELETE("/:id", r.Controller.User.Delete)

	// Product routes
	product.POST("/", r.Controller.Product.Create)
	product.GET("/", r.Controller.Product.GetAll)
	product.GET("/:id", r.Controller.Product.GetByID)
	product.PUT("/:id", r.Controller.Product.Update)
	product.DELETE("/:id", r.Controller.Product.Delete)

	// Category routes
	category.POST("/", r.Controller.Category.Create)
	category.GET("/", r.Controller.Category.GetAll)
	category.GET("/:id", r.Controller.Category.GetByID)
	category.PUT("/:id", r.Controller.Category.Update)
	category.DELETE("/:id", r.Controller.Category.Delete)

	return router
}
