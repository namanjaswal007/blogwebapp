package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func ConnectDB() (db *gorm.DB, err error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=blog_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}
