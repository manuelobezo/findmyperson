package data

import (
	"context"
	"time"
	//"github.com/manuelobezo/go-postgres-ambertAlert/pkg/person"
	"github.com/manuelobezo/go-postgres-ambertAlert/pkg/found"
    "fmt"
)

type FoundRepository struct {
    Data *Data
}

//insertar reporte de persona encontrada en la tabla
func (fr *FoundRepository) Create(ctx context.Context, f *found.Found) error {
	error  :=  fr.Data.DB.QueryRow("SELECT id FROM persons WHERE curp = $1;",f.Person.Curp).Scan(&f.PersonID) 

	if error != nil{
		id :=fr.Data.DB.QueryRow("INSERT INTO persons (first_name, last_name, curp, birthdate, last_seen, missing, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;", f.Person.FirstName, f.Person.LastName, f.Person.Curp,f.Person.BirthDate, "", false,time.Now(),time.Now()).Scan(&f.PersonID) 
		if id != nil {
			return id
		}
	}
	f.Person.ID=f.PersonID
	//Update missing to false
	fr.Data.DB.QueryRow("UPDATE persons SET missing=false WHERE id = $1;",f.PersonID)
	
	q := `
    INSERT INTO founds (person_id, found_date, address, hospitalized, condition, more_info,  created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id;
	`
	fmt.Println(f.Person)
    current_time := time.Now()
    f.CreatedAt = current_time
    f.UpdatedAt = current_time

    row := fr.Data.DB.QueryRowContext(
        ctx, q, f.PersonID, f.FoundDate, f.Address, f.Hospitalized, f.Condition, f.MoreInfo, f.CreatedAt,
         f.UpdatedAt,
    )

    err := row.Scan(&f.ID)
    if err != nil {
        return err
    }

    return nil
}

//obtener todos
func (fr *FoundRepository) GetAll(ctx context.Context) ([]found.Found, error) {
    q := `
    SELECT id, person_id, found_date, address, hospitalized, condition, more_info, created_at, updated_at
        FROM founds;
    `

    rows, err := fr.Data.DB.QueryContext(ctx, q)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var founds []found.Found
    for rows.Next() {
        var f found.Found
        rows.Scan(&f.ID, &f.PersonID, &f.FoundDate, &f.Address, &f.Hospitalized, &f.Condition, &f.MoreInfo, &f.CreatedAt,&f.UpdatedAt)
		er  := fr.Data.DB.QueryRow("SELECT id, first_name, last_name, curp, birthdate FROM persons WHERE id = $1;",f.PersonID).Scan(&f.Person.ID,&f.Person.FirstName,&f.Person.LastName, &f.Person.Curp, &f.Person.BirthDate ) 
		if er != nil {
			return nil, er
		}

		f.Person.BirthDate = f.Person.BirthDate[0:10]//format to yyyy-mm-dd

        founds = append(founds, f)
    }

    return founds, nil
}