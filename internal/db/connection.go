package db

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/tuki277/golang-boilerplate/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.DBConfig) (*gorm.DB, error) {
	log.Info(dsn(cfg))
	db, err := gorm.Open(postgres.Open(dsn(cfg)), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open db connection: %w", err)
	}

	return db, nil
}

func dsn(c config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name,
	)
}
