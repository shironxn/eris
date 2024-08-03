package main

import (
	"fmt"

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

	fmt.Println("SERVER RUNNING")
	if err := config.NewServer(cfg.Server, db).Run(); err != nil {
		panic(err)
	}
}
