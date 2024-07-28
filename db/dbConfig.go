package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) (*gorm.DB, error) {

	// fmt.Println("Start of database File  ")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// 	fmt.Println("Database connection Successfully   ", DB)
	if err != nil {
		fmt.Println("Error in server :", err)
	}

	return DB, err
}
