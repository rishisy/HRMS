package main

import (
	"HRMS/db"
	"HRMS/doctors"
	"HRMS/patients"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	fmt.Println(db.DB)
	fmt.Println("Main file started with database ", db.DB)
	// main exposed point for the application
	router := gin.Default()

	router.POST("/patient", patients.PostPatient)
	router.GET("/patient/:id", patients.GetPatientByID)
	router.GET("/fetchPatientByDoctorId/:doctorId", patients.GetPatientByDoctorID)
	router.PATCH("patient/:id", patients.PatchPatient)

	router.POST("/doctor", doctors.PostDoctor)
	router.GET("/doctor/:id", doctors.GetDoctorByID)
	router.PATCH("doctor/:id", doctors.PatchDoctor)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

}
