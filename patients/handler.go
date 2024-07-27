package patients

import (
	"HRMS/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

// Validator object for validation of struct fields
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// PostPatient adds an album from JSON received in the request body.
func PostPatient(c *gin.Context) {

	var newPatient Patient

	// Call BindJSON to bind the received JSON to newPatients
	if err := c.BindJSON(&newPatient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation checks for struct fields
	if err := validate.Struct(newPatient); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	// Fill extra fields
	newPatient.ID = db.GenerateUUID()
	newPatient.CreatedAt = time.Now()
	newPatient.UpdatedAt = time.Now()

	// Save the new album to the database
	if err := db.DB.Create(&newPatient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newPatient)
}

func GetPatientByID(c *gin.Context) {

	ID := c.Param("id")

	if len(ID) != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": "Invalid Id"})
		return
	}

	var newPatient Patient
	newPatient.ID = ID

	if result := db.DB.First(&newPatient); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newPatient)

}

func PatchPatient(c *gin.Context) {

	ID := c.Param("id")

	if len(ID) != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": "Invalid Id"})
		return
	}

	var patch PatientPatchObject

	// Call BindJSON to bind the received JSON to Patch Object
	if err := c.BindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if at least one field is provided
	if patch.ContactNo == nil && patch.Address == nil && patch.DoctorId == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "At least one field must be provided"})
		return
	}
	// validation checks for struct fields
	if err := validate.Struct(patch); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	var patient Patient

	// Receive the Original patient from the database
	if result := db.DB.First(&patient); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": result.Error.Error()})
		return
	}

	if patch.ContactNo != nil {
		patient.ContactNo = *patch.ContactNo
	}

	if patch.DoctorId != nil {
		patient.DoctorId = *patch.DoctorId
	}
	if patch.Address != nil {
		patient.Address = *patch.Address
	}
	patient.UpdatedAt = time.Now()

	if result := db.DB.Save(&patient); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, patient)

}

func GetPatientByDoctorID(c *gin.Context) {

	doctorID := c.Param("doctorId")

	if len(doctorID) != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": "Invalid Doctor Id"})
		return
	}

	var patients []Patient

	if result := db.DB.Where("doctor_id = ?", doctorID).Find(&patients); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if len(patients) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"Error": "No Patients for given Doctor Id "})
	}

	c.IndentedJSON(http.StatusOK, patients)

}
