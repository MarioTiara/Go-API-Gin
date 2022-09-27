package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Mario2022! dbname=pustaka-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
