package db

import (
	"fmt"
	"kienmatu/go-todos/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(cfg *config.Configuration) *gorm.DB {
	//dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	//db.AutoMigrate(&models.Person{})
	return db
}
