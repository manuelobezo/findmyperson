package data

import (
	"context"
	"time"
	"github.com/manuelobezo/go-postgres-ambertAlert/pkg/person"
)

type PersonRepository struct {
    Data *Data
}

//obtener todos
func (pr *PersonRepository) GetAll(ctx context.Context) ([]person.Person, error) {
    q := `
    SELECT id, first_name, last_name, curp, found, birthdate
        created_at, updated_at
        FROM persons;
    `

    rows, err := pr.Data.DB.QueryContext(ctx, q)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var persons []person.Person
    for rows.Next() {
        var p person.Person
        rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Curp, &p.BirthDate, &p.Found,
			&p.CreatedAt, &p.UpdatedAt)
        persons = append(persons, p)
    }

    return persons, nil
}

//obtener por id
func (pr *PersonRepository) GetOne(ctx context.Context, id uint) (person.Person, error) {
    q := `
    SELECT id, first_name, last_name, curp, found, birthdate
        created_at, updated_at
        FROM persons WHERE id = $1;
    `

    row := pr.Data.DB.QueryRowContext(ctx, q, id)

    var p person.Person
    err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Curp, &p.BirthDate, &p.Found,
		&p.CreatedAt, &p.UpdatedAt)
    if err != nil {
        return person.Person{}, err
    }

    return p, nil
}

//Obtener por curp
func (pr *PersonRepository) GetByCurp(ctx context.Context, curp string) (person.Person, error) {
    q := `
    SELECT id, first_name, last_name, curp, found, birthdate
        created_at, updated_at
        FROM persons where curp = $1;
    `

    row := pr.Data.DB.QueryRowContext(ctx, q, curp)

    var p person.Person
    err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Curp, &p.BirthDate, &p.Found,
		&p.CreatedAt, &p.UpdatedAt)
    if err != nil {
        return person.Person{}, err
    }

    return p, nil
}

//insertar persona en la tabla
func (pr *PersonRepository) Create(ctx context.Context, p *person.Person) error {
    q := `
    INSERT INTO persons (first_name, last_name, curp, birthdate, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `


    row := pr.Data.DB.QueryRowContext(
        ctx, q, p.FirstName, p.LastName, p.Curp, p.BirthDate, time.Now(),
         time.Now(),
    )

    err := row.Scan(&p.ID)
    if err != nil {
        return err
    }

    return nil
}

//actualizar persona
func (pr *PersonRepository) Update(ctx context.Context, id uint, p person.Person) error {
    q := `
    UPDATE persons set first_name=$1, last_name=$2, curp=$3, found=$4, birthdate=$5, updated_at=$6
        WHERE id=$7;
    `

    stmt, err := pr.Data.DB.PrepareContext(ctx, q)
    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.ExecContext(
        ctx, p.FirstName, p.LastName, p.Curp,
        p.Found, p.BirthDate, time.Now(), id,
    )
    if err != nil {
        return err
    }

    return nil
}

//Eliminar persona
func (pr *PersonRepository) Delete(ctx context.Context, id uint) error {
    q := `DELETE FROM users WHERE id=$1;`

    stmt, err := pr.Data.DB.PrepareContext(ctx, q)
    if err != nil {
        return err
    }

    defer stmt.Close()

    _, err = stmt.ExecContext(ctx, id)
    if err != nil {
        return err
    }

    return nil
}