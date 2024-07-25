package main

import (
	"HRMS/db"
	"HRMS/doctors"
	"HRMS/patients"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
}

func main() {
	DB := db.Connect()
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&doctors.Doctor{})
	DB.AutoMigrate(&patients.Patient{})

	var userrandom User
	userrandom.ID = 12
	userrandom.Name = db.GenerateUUID()
	errr := DB.Create(&userrandom)
	print(errr)

	var user2 []User
	DB.Find(&user2)

	for _, val := range user2 {
		fmt.Println(val.Name)
	}

}
