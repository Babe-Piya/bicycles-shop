package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB : Main database instance
var DB *gorm.DB

// ConnectDB : Initial connect database function
func ConnectDB() error {
	dbAddress := "localhost"
	dbUser := "postgres"
	dbPort := "5432"
	dbPass := "example"
	dbName := "postgres"
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbAddress,
		dbPort,
		dbUser,
		dbPass,
		dbName)

	var err error

	DB, err = gorm.Open("postgres", connectionString)

	if err != nil {
		return err
	}

	//DB.DB().SetMaxIdleConns(5)
	//DB.DB().SetMaxOpenConns(5)
	//DB.DB().SetConnMaxLifetime(time.Minute * 5)

	return err
}
