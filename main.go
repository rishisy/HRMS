package main

import (
	"HRMS/db"
	"fmt"
)

func main() {

	fmt.Println("Start of main file")
	DB := db.Connect()
	fmt.Println("main file has access to ", DB)

}
