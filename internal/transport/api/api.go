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

func (a *API) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/url", a.CreateURL()).Methods(http.MethodPost)
	r.HandleFunc("/v1/url", a.GetURL()).Methods(http.MethodGet)
	return r
}

func (a *API) CreateURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (a *API) GetURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
