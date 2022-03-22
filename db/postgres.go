package db

import (
	"fmt"
	"kienmatu/go-todos/config"
	"kienmatu/go-todos/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(cfg *config.Configuration, migrate bool) *gorm.DB {
	//dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if migrate {
		db.AutoMigrate(&models.User{}, &models.Todo{})
		// if err != nil {
		// 	panic("Error when run migrations")
		// }
	}
	return db
}
