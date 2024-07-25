package patients

import "time"

type Patient struct {
	ID        string    `json:"id"` //golang uses ID filed as primary key by default
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	ContactNo string    `json:"contact_no"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DoctorId  string    `json:"doctor_id"` // foreign key
}
