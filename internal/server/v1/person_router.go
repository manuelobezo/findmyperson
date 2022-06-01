package v1

import( 

		"encoding/json"
		"fmt"
		"net/http"
        "strconv"
        //"time"

		"github.com/go-chi/chi"
		"github.com/manuelobezo/go-postgres-ambertAlert/pkg/response"
        "github.com/manuelobezo/go-postgres-ambertAlert/pkg/person"
        
        //"github.com/mitchellh/mapstructure"

)
type PersonRouter struct {
    Repository person.Repository
}

//crear persona
func (pr *PersonRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
    /*var p person.Person
    //err := json.NewDecoder(r.Body).Decode(&p)

    decoder := json.NewDecoder(r.Body)

    // Se usa para almacenar datos de clave de parámetro = valor
    var params map[string]string
 
    // Analiza los parámetros en el mapa
    err := decoder.Decode(&params)

    const shortForm = "2006-01-02"
    t, _ := time.Parse(shortForm, params["BirthDate"])
    

    mapstructure.Decode(params, &p)
    p.BirthDate=t
    fmt.Printf("%+v\n",p)

    */
    var p person.Person
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    defer r.Body.Close()

    ctx := r.Context()
    err = pr.Repository.Create(ctx, &p)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
    response.JSON(w, r, http.StatusCreated, response.Map{"person": p})
    
}

//obtener personas
func (pr *PersonRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    persons, err := pr.Repository.GetAll(ctx)
    if err != nil {
        response.HTTPError(w, r, http.StatusNotFound, err.Error())
        return
    }

    response.JSON(w, r, http.StatusOK, response.Map{"persons": persons})
}

//Obtener una persona
func (pr *PersonRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    ctx := r.Context()
    p, err := pr.Repository.GetOne(ctx, uint(id))
    if err != nil {
        response.HTTPError(w, r, http.StatusNotFound, err.Error())
        return
    }

    response.JSON(w, r, http.StatusOK, response.Map{"person": p})
}

//obtener por cupr
func (pr *PersonRouter) GetOneHandlerCurp(w http.ResponseWriter, r *http.Request) {
    curp := chi.URLParam(r, "curp")


    ctx := r.Context()
    p, err := pr.Repository.GetByCurp(ctx, string(curp))
    if err != nil {
        response.HTTPError(w, r, http.StatusNotFound, err.Error())
        return
    }

    response.JSON(w, r, http.StatusOK, response.Map{"person": p})
}

//Actualizar persona
func (pr *PersonRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    var p person.Person
    err = json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    defer r.Body.Close()

    ctx := r.Context()
    err = pr.Repository.Update(ctx, uint(id), p)
    if err != nil {
        response.HTTPError(w, r, http.StatusNotFound, err.Error())
        return
    }

    response.JSON(w, r, http.StatusOK, nil)
}

//borrar personsa
func (pr *PersonRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        response.HTTPError(w, r, http.StatusBadRequest, err.Error())
        return
    }

    ctx := r.Context()
    err = pr.Repository.Delete(ctx, uint(id))
    if err != nil {
        response.HTTPError(w, r, http.StatusNotFound, err.Error())
        return
    }

    response.JSON(w, r, http.StatusOK, response.Map{})
}

//router

func (pr *PersonRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", pr.GetAllHandler)

	r.Post("/", pr.CreateHandler)

	r.Get("/{id}", pr.GetOneHandler)

	//r.Get("/{curp}", pr.GetOneHandlerCurp)

	r.Put("/{id}", pr.UpdateHandler)

	r.Delete("/{id}", pr.DeleteHandler)

	return r
}