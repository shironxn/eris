package config

import (
	"net/http"

	"github.com/shironxn/eris/internal/app/controller"
	"github.com/shironxn/eris/internal/app/model"
	https "github.com/shironxn/eris/internal/infrastructure/http"
	"github.com/shironxn/eris/internal/infrastructure/repository"
	"github.com/shironxn/eris/internal/infrastructure/service"
	"github.com/shironxn/eris/internal/infrastructure/util"
	"gorm.io/gorm"
)

type Server struct {
	Host   string
	Port   string
	Router https.Router
	DB     *gorm.DB
	JWT    util.JWT
}

func NewServer(server Server) *Server {
	return &Server{
		Host: server.Host,
		Port: server.Port,
		DB:   server.DB,
		JWT:  server.JWT,
	}
}

func (s *Server) Run() error {
	jwt := util.NewJWT(s.JWT)

	middleware := https.NewMiddleware(*jwt)

	userRepository := repository.NewUserRepository(s.DB)
	productRepository := repository.NewProductRepository(s.DB)
	categoryRepository := repository.NewCategoryRepository(s.DB)

	userService := service.NewUserService(userRepository)
	productService := service.NewProductService(productRepository)
	categoryService := service.NewCategoryService(categoryRepository)

	userController := controller.NewUserController(userService, *jwt)
	productController := controller.NewProducController(productService)
	categoryController := controller.NewCategoryController(categoryService)

	router := https.NewRouter(
		https.Controllers{
			User:     userController,
			Product:  productController,
			Category: categoryController,
		},
		https.Middleware{
			JWT: middleware.JWT,
		},
	)

	s.DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{})

	server := &http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router.Route(),
	}

	return server.ListenAndServe()
}
