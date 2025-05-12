package helpers

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm(dbMock *sql.DB) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbMock,
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
