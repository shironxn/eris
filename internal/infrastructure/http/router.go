package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shironxn/eris/internal/app/controller"
)

type Router struct {
	User     controller.UserController
	Product  controller.ProductController
	Category controller.CategoryController
}

func NewRouter(r Router) *Router {
	return &Router{
		User:     r.User,
		Product:  r.Product,
		Category: r.Category,
	}
}

func (r *Router) Route() http.Handler {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	user := v1.Group("/users")
	product := v1.Group("/products")
	category := v1.Group("/categories")

	auth.POST("/login", r.User.Login)
	auth.POST("/register", r.User.Register)

	user.GET("/", r.User.GetAll)
	user.GET("/:id", r.User.GetByID)
	user.PUT("/:id", r.User.Update)
	user.DELETE("/:id", r.User.Delete)

	product.POST("/", r.Product.Create)
	product.GET("/", r.Product.GetAll)
	product.GET("/:id", r.Product.GetByID)
	product.PUT("/:id", r.Product.Update)
	product.DELETE("/:id", r.Product.Delete)

	category.POST("/", r.Category.Create)
	category.GET("/", r.Category.GetAll)
	category.GET("/:id", r.Category.GetByID)
	category.PUT("/:id", r.Category.Update)
	category.DELETE("/:id", r.Category.Delete)

	return router
}
