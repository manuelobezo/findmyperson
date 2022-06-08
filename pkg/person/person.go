package person

import "time"

type Person struct {
    ID           uint      `json:"id,omitempty"`
    FirstName    string    `json:"first_name,omitempty"`
    LastName     string    `json:"last_name,omitempty"`
	Curp     	 string    `json:"curp,omitempty"`
    BirthDate    string    `json:"birth_date,omitempty"`	
    LastSeen     string    `json:"last_seen,omitempty"`
    Missing		 bool	   `json:"missing,omitempty"`	 
    CreatedAt    time.Time `json:"created_at,omitempty"`
    UpdatedAt    time.Time `json:"updated_at,omitempty"`
}