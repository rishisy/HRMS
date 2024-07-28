package main

import (
	"HRMS/doctors"
	"HRMS/patients"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server for creating a single point for whole app to access server objects
type Server struct {
	DB *gorm.DB
	// logs
	// more server objects
}

func NewServer(db *gorm.DB) *Server {
	return &Server{DB: db}
}

// Start initializes the HTTP server
func (s *Server) Start(addr string) error {
	router := gin.Default()

	router.POST("/patient", patients.PostPatient)
	router.GET("/patient/:id", patients.GetPatientByID)
	router.GET("/fetchPatientByDoctorId/:doctorId", patients.GetPatientByDoctorID)
	router.PATCH("patient/:id", patients.PatchPatient)

	router.POST("/doctor", doctors.PostDoctor)
	router.GET("/doctor/:id", doctors.GetDoctorByID)
	router.PATCH("doctor/:id", doctors.PatchDoctor)

	return router.Run(addr)
}
