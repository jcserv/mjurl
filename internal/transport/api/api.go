package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcserv/mjurl/internal/url"
	"github.com/jcserv/mjurl/model"
)

const (
	APIV1URLPath = "/api/v1/url"
	GetURLPath   = APIV1URLPath + "/{short}"
)

type API struct {
	URLService model.IURLService
}

func NewAPI(urlService model.IURLService) *API {
	return &API{
		URLService: urlService,
	}
}

func (a *API) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(APIV1URLPath, a.CreateURL()).Methods(http.MethodPost)
	r.HandleFunc(GetURLPath, a.GetURL()).Methods(http.MethodGet)
	return r
}

func (a *API) CreateURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (a *API) GetURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		cmd, err := url.NewGetURLByShort(vars["short"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := cmd.Execute(ctx, a.URLService)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		obj, err := json.Marshal(u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(obj)
	}
}
