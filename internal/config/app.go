package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Server   Server
	Database Database
}

func New() (*App, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	app := App{
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
	}

	return &app, nil
}
