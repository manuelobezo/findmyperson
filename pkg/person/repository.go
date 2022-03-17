package person

import "context"

// Repository handle the CRUD operations with Person.
type Repository interface {
    GetAll(ctx context.Context) ([]Person, error)
    GetOne(ctx context.Context, id uint) (Person, error)
    GetByCurp(ctx context.Context, curp string) (Person, error)
    Create(ctx context.Context, person *Person) error
    Update(ctx context.Context, id uint, person Person) error
    Delete(ctx context.Context, id uint) error
}