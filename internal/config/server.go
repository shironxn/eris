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
	JWT    util.JWT
	Router https.Router
	DB     *gorm.DB
}

func NewServer(server Server, db *gorm.DB) *Server {
	return &Server{
		Host:   server.Host,
		Port:   server.Port,
		JWT:    server.JWT,
		Router: server.Router,
		DB:     db,
	}
}

func (s *Server) Run() error {
	jwt := util.NewJWT(s.JWT)
	middleware := https.NewMiddleware(*jwt)

	// Initializations
	userRepository := repository.NewUserRepository(s.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository(s.DB)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProducController(productService)

	categoryRepository := repository.NewCategoryRepository(s.DB)
	categoryService := service.NewCategoryService(categoryRepository)
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
