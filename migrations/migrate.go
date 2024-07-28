package main

import (
	"HRMS/db"
	"HRMS/doctors"
	"HRMS/patients"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// for getting all environment variables with error checks
func getEnvVar(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	return os.Getenv(key)
}

func main() {

	// Database connections ( personally using postgres in this case )
	DB, err := db.Connect(getEnvVar("DB_STRING"))
	if err != nil {
		fmt.Println("Connection to Database Failed") //log can be used instead of println for maintaining logfiles
		return
	}
	//Migration : Only Uncomment if you  change the patients/doctors Models

	if err := DB.AutoMigrate(&doctors.Doctor{}); err != nil {
		fmt.Println("Migration falied for doctor ")
		return
	}
	if err := DB.AutoMigrate(&patients.Patient{}); err != nil {
		fmt.Println("Migration falied for Patient")
		return
	}

	fmt.Println("Migration Performed Successfully")

}
