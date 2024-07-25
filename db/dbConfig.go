package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	fmt.Println("Start of database File  ")
	dsn := "?"
	DB, dberr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dberr != nil {
		panic(dberr)
	}
	fmt.Println("Database connection Successfull   ", DB)

	return DB

}
