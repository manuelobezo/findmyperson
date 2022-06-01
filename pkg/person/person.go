package person

import "time"

type Person struct {
    ID           uint      `json:"id,omitempty"`
    FirstName    string    `json:"first_name,omitempty"`
    LastName     string    `json:"last_name,omitempty"`
	Curp     	 string    `json:"curp,omitempty"`
	Found		 bool	   `json:"found,omitempty"`
	BirthDate    string    `json:"birth_date,omitempty"`		 
    CreatedAt    time.Time `json:"created_at,omitempty"`
    UpdatedAt    time.Time `json:"updated_at,omitempty"`
}