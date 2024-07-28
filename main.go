package main

import (
	"HRMS/db"
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

	// Starting Server by creating a server object common across all files
	server := NewServer(DB)
	if err := server.Start(":" + getEnvVar("PORT")); err != nil {
		fmt.Println("SERVER ERROR ", err) //log can be used instead of println for maintaining logfiles
		return
	}

	fmt.Println("Server running Successfully on ", getEnvVar("PORT"))
}
