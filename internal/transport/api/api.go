package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	APIV1URLPath = "/api/v1/url"
)

type API struct {
}

func NewAPI() *API {
	return &API{}
}

func (a *API) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(APIV1URLPath, a.CreateURL()).Methods(http.MethodPost)
	r.HandleFunc(APIV1URLPath, a.GetURL()).Methods(http.MethodGet)
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
