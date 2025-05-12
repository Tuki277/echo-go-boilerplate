package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tuki277/golang-boilerplate/internal/config"
	"gorm.io/gorm"
)

type Server struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(
	echo *echo.Echo,
	db *gorm.DB,
	config *config.Config,
) *Server {
	return &Server{
		Echo:   echo,
		DB:     db,
		Config: config,
	}
}

func (s *Server) Start(addr string) error {
	if err := s.Echo.Start(":" + addr); err != nil && errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("start echo: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.Echo.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown echo: %w", err)
	}

	return nil
}
