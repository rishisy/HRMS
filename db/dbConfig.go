package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	fmt.Println("Start of database File  ")

	dsn := "user=postgres.yxdsfhdvljmvkiewsfcx password=doctor@44rishi host=aws-0-ap-south-1.pooler.supabase.com port=6543 dbname=postgres"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Unable to connect to PostgreSQL Database")
		panic(err)
	}

	fmt.Println("Database connection Successfull   ", DB)

	return DB

}
