package helpers

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/labstack/echo/v4"
	"github.com/tuki277/golang-boilerplate/internal/config"
	"github.com/tuki277/golang-boilerplate/internal/server"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *server.Server {
	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Failed to parse envs:", err.Error())
	}

	s := &server.Server{
		Echo:   echo.New(),
		DB:     db,
		Config: &cfg,
	}

	return s
}
