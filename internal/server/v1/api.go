package v1

import (
    "net/http"

	"github.com/go-chi/chi"
	
	"github.com/manuelobezo/go-postgres-ambertAlert/internal/data"
)

func New() http.Handler {
    r := chi.NewRouter()


	pr := &PersonRouter{
        Repository: &data.PersonRepository{
            Data: data.New(),
        },
	}
	
	r.Mount("/persons", pr.Routes())
    return r
}