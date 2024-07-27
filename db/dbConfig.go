package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	fmt.Println("Start of database File  ")
	
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection Successfull   ", DB)

	// Migration : Only Uncomment after changes made to patients/doctors Models
	//DB.AutoMigrate(&User{})
	//DB.AutoMigrate(&doctors.Doctor{})
	//DB.AutoMigrate(&patients.Patient{})

}
