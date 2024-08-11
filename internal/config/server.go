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
	jwt    util.JWT
	router https.Router
	db     *gorm.DB
}

func NewServer(server Server, db *gorm.DB) *Server {
	return &Server{
		Host:   server.Host,
		Port:   server.Port,
		jwt:    server.jwt,
		router: server.router,
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

	jwt := util.NewJWT(s.jwt)

	middleware := https.NewMiddleware(jwt)

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

	s.db.AutoMigrate(&model.User{}, &model.Product{}, &model.Category{})
	server := &http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router.Route(),
	}

	return server.ListenAndServe()
}
