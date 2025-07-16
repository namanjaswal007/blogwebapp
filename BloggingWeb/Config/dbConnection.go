package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	models "BloggingWeb/Model"

)

// func ConnectDB(dbName string) (database models.Database, err error) {
// 	connStr := "host=localhost port=5432 user=postgres password=1730 dbname=" + dbName + " sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to open DB: %v", err))
// 	}
// 	db.AutoMigrate(&view.User{}, &view.Blog{})
// 	database = models.Database{MainDB: db}
// 	return
// }

func ConnectDB(dbName string) (database models.Database, err error) {
	connStr := fmt.Sprintf("host=localhost port=5432 user=postgres password=1730 dbname=%s sslmode=disable", dbName)

	// Open DB connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return database, fmt.Errorf("failed to open DB: %w", err)
	}

	// Test the DB connection
	if err = db.Ping(); err != nil {
		return database, fmt.Errorf("failed to ping DB: %w", err)
	}

	// Return custom Database struct
	database = models.Database{MainDB: db}
	fmt.Println("Connected to PostgreSQL with sql package.")
	return
}
