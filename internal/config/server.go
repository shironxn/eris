package config

import (
	"github.com/shironxn/eris/internal/app/controller"
	"github.com/shironxn/eris/internal/app/model"
	https "github.com/shironxn/eris/internal/infrastructure/http"
	"github.com/shironxn/eris/internal/infrastructure/repository"
	"github.com/shironxn/eris/internal/infrastructure/service"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	Host   string
	Port   string
	Router https.Router
	db     *gorm.DB
}

func NewServer(cfg Server, db *gorm.DB) *Server {
	return &Server{
		Host:   cfg.Host,
		Port:   cfg.Port,
		Router: cfg.Router,
		db:     db,
	}
}

func (s *Server) Run() error {
	userRepository := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepository(s.db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProducController(productService)

	categoryRepository := repository.NewCategoryRepository(s.db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	router := https.NewRouter(https.Router{
		User:     userController,
		Product:  productController,
		Category: categoryController,
	})

	s.db.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{})
	server := &http.Server{
		Addr:    ":" + s.Port,
		Handler: router.Route(),
	}

	return server.ListenAndServe()
}
