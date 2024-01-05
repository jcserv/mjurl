package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func (a *API) GetURL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (a *API) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", a.GetURL).Methods(http.MethodGet)
	return r
}
