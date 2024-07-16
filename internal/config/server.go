package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	https "github.com/shironxn/eris/internal/infrastructure/http"
)

type Server struct {
	Host   string
	Port   string
	Router https.Router
}

func NewServer(cfg Server) *Server {
	return &Server{
		Host:   cfg.Host,
		Port:   cfg.Port,
		Router: cfg.Router,
	}
}

func (s *Server) Run() error {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "HI",
		})
	})

	// users := router.Group("/api/v1/users")
	// {
	// }

	server := &http.Server{
		Addr:    ":" + s.Port,
		Handler: s.Router.Route(),
	}

	return server.ListenAndServe()
}
