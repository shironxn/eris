package main

import (
	"github.com/shironxn/eris/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := config.NewDatabase(cfg.Database).Connection()
	if err != nil {
		panic(err)
	}

	if err := config.NewServer(config.Server{
		Host: cfg.Server.Host,
		Port: cfg.Server.Port,
		JWT:  cfg.Server.JWT,
		DB:   db,
	}).Run(); err != nil {
		panic(err)
	}
}
