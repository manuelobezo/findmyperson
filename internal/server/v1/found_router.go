package v1

import( 

		"encoding/json"
		"fmt"
		"net/http"
        //"strconv"
        //"time"

		"github.com/go-chi/chi"
		"github.com/manuelobezo/go-postgres-ambertAlert/pkg/response"
		//"github.com/manuelobezo/go-postgres-ambertAlert/pkg/person"
		"github.com/manuelobezo/go-postgres-ambertAlert/pkg/found"
        
        //"github.com/mitchellh/mapstructure"

)

type FoundRouter struct {
    Repository found.Repository
}


//crear reporte encontrado
func (fr *FoundRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {

    var f found.Found
    err := json.NewDecoder(r.Body).Decode(&f)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    defer r.Body.Close()

    ctx := r.Context()
    err = fr.Repository.Create(ctx, &f)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), f.ID))
    response.JSON(w, r, http.StatusCreated, response.Map{"found": f})
    
}

//obtener reporte de personas encontradas
func (fr *FoundRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    founds, err := fr.Repository.GetAll(ctx)
    if err != nil {
        response.HTTPError(w, r, http.StatusNotFound, err.Error())
        return
    }

    response.JSON(w, r, http.StatusOK, response.Map{"founds": founds})
}

//router

func (fr *FoundRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", fr.GetAllHandler)

	r.Post("/", fr.CreateHandler)

	//r.Get("/{id}", pr.GetOneHandler)

	//r.Get("/{curp}", pr.GetOneHandlerCurp)

	//r.Put("/{id}", pr.UpdateHandler)

	//r.Delete("/{id}", pr.DeleteHandler)

	return r
}