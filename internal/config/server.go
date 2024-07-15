package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Host string
	Port string
}

func NewServer(cfg Server) *Server {
	return &Server{
		Host: cfg.Host,
		Port: cfg.Port,
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
		Handler: router,
	}

	return server.ListenAndServe()
}
