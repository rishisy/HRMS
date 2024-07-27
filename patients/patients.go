package patients

import "time"

type Patient struct {
	ID        string    `json:"id" gorm:"primaryKey"` //golang uses ID filed as primary key by default
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Address   string    `json:"address" validate:"required,min=3,max=100"`
	ContactNo string    `json:"contact_no" validate:"required,len=10"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DoctorId  string    `json:"doctor_id" validate:"required,len=5"` // foreign key
}
