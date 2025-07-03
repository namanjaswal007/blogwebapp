package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	models "BloggingWeb/Model"
	view "BloggingWeb/View"

)

func ConnectDB(dbName string) (database models.Database, err error) {
	connStr := "host=localhost port=5432 user=postgres password=1730 dbname="+dbName+" sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to open DB: %v", err))
	}
	db.AutoMigrate(&view.User{}, &view.Blog{})
	database = models.Database{MainDB: db}
	return
}
