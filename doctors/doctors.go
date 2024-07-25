package doctors

import "time"

type Doctor struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ContactNo string    `json:"contact_no"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
