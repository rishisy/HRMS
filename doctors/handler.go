package doctors

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

// PostDoctor adds an album from JSON received in the request body.
func PostDoctor(c *gin.Context) {

	var newDoctor Doctor

	// Call BindJSON to bind the received JSON to newDoctors
	if err := c.BindJSON(&newDoctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation checks for struct fields
	if err := validate.Struct(newDoctor); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": validationErrors})
		return
	}

	// Fill extra fields
	newDoctor.ID = db.GenerateUUID()
	newDoctor.CreatedAt = time.Now()
	newDoctor.UpdatedAt = time.Now()

	// Save the new album to the database
	if err := db.DB.Create(&newDoctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newDoctor)
}

func GetDoctorByID(c *gin.Context) {

	ID := c.Param("id")

	if len(ID) != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": "Invalid Id"})
		return
	}

	var newDoctor Doctor
	newDoctor.ID = ID

	if result := db.DB.First(&newDoctor); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, newDoctor)

}

func PatchDoctor(c *gin.Context) {

	ID := c.Param("id")

	if len(ID) != 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": "Invalid Id"})
		return
	}

	var patch DoctorPatchObject

	// Call BindJSON to bind the received JSON to Patch Object
	if err := c.BindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	var doctor Doctor

	// Receive the Original doctor from the database
	if result := db.DB.First(&doctor); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error :": result.Error.Error()})
		return
	}

	if patch.ContactNo != nil {
		doctor.ContactNo = *patch.ContactNo
	}
	doctor.UpdatedAt = time.Now()

	if result := db.DB.Save(&doctor); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, doctor)

}
