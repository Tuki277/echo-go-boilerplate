package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/tuki277/golang-boilerplate/docs"
	"github.com/tuki277/golang-boilerplate/internal/config"
	"github.com/tuki277/golang-boilerplate/internal/db"
	"github.com/tuki277/golang-boilerplate/internal/server"
	"github.com/tuki277/golang-boilerplate/internal/server/routes"

	"github.com/caarlos0/env/v11"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const shutdownTimeout = 20 * time.Second

//	@title			Echo Demo App
//	@version		1.0
//	@description	This is a demo version of Echo app.

//	@contact.name	HuyTH (CTO - Cheif TrongXe Office - iKame Global)
//	@contact.url	https://huybeos2707@gmail.com - huyth@ikameglobal.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

// @BasePath	/
func main() {
	if err := run(); err != nil {
		slog.Error("Service run error", "err", err.Error())
		os.Exit(1)
	}
}

func run() error {
	envFlag := flag.String("ENV", "env", "Environment file to load (e.g., env, env.production)")
	flag.Parse()

	var envFile string
	switch *envFlag {
	case "env.production":
		envFile = ".env.production"
	default:
		envFile = ".env"
	}

	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("load env file %s: %w", envFile, err)
	}

	var cfg config.Config
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("parse env: %w", err)
	}

	log.Info(cfg.HTTP.Port)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	gormDB, err := db.NewGormDB(cfg.DB)
	if err != nil {
		return fmt.Errorf("new db connection: %w", err)
	}

	app := server.NewServer(echo.New(), gormDB, &cfg)

	routes.ConfigureRoutes(app)

	go func() {
		if err = app.Start(cfg.HTTP.Port); err != nil {
			slog.Error("Server error", "err", err.Error())
		}
	}()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	<-shutdownChannel

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("http server shutdown: %w", err)
	}

	dbConnection, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("get db connection: %w", err)
	}

	if err := dbConnection.Close(); err != nil {
		return fmt.Errorf("close db connection: %w", err)
	}

	return nil
}
