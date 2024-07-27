package doctors

import "time"

type Doctor struct {
	ID        string    `json:"id" gorm:"primaryKey"` //golang uses ID filed as primary key by default
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	ContactNo string    `json:"contact_no" validate:"required,len=10"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DoctorPatchObject struct {
	ContactNo *string `json:"contact_no" validate:"required,len=10"`
}
