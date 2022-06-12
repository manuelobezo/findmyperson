package found

import (
	"time"
)

type Person struct {
	ID           uint      `json:"id,omitempty"`
    FirstName    string    `json:"first_name,omitempty"`
    LastName     string    `json:"last_name,omitempty"`
	Curp     	 string    `json:"curp,omitempty"`
    BirthDate    string    `json:"birth_date,omitempty"`	 
}

type Found struct {
    ID           uint      `json:"id,omitempty"`
    PersonID     uint	   `json:"person_id,omitempty"`
    FoundDate    string    `json:"found_date,omitempty"`
	Address      string    `json:"address,omitempty"`
    Hospitalized string    `json:"hospitalized,omitempty"`	
    Condition    string    `json:"condition,omitempty"`
	MoreInfo	 string	   `json:"more_info,omitempty"`	 
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Person 		 Person    `json:"person,omitempty"`
}