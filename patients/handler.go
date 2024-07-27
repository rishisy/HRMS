package patients

import (
	"HRMS/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

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

func GetPatient(c *gin.Context) {

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
