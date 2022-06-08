package found

import "context"

// Repository handle the CRUD operations with Person.
type Repository interface {
	Create(ctx context.Context, found *Found) error
    GetAll(ctx context.Context) ([]Found, error)
    //GetOne(ctx context.Context, id uint) (Person, error)
    //GetByCurp(ctx context.Context, curp string) (Person, error)
    
    //Update(ctx context.Context, id uint, person Person) error
    //Delete(ctx context.Context, id uint) error
}