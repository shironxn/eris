package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/shironxn/eris/internal/infrastructure/util"
)

type App struct {
	Server   Server
	Database Database
	JWT      util.JWT
}

func New() (*App, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &App{
		Server: Server{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Port: os.Getenv("DB_PORT"),
		},
		JWT: util.JWT{
			Access:  os.Getenv("JWT_ACCESS"),
			Refresh: os.Getenv("JWT_REFRESH"),
		},
	}, nil
}
