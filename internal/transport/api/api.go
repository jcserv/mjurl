package api

import (
	"github.com/gorilla/mux"
	v1 "github.com/jcserv/mjurl/internal/transport/api/v1"
	"github.com/jcserv/mjurl/model"
)

const (
	APIV1URLPath = "/api/v1/url"
	GetURLPath   = APIV1URLPath + "/{short}"
)

type API struct {
	V1API *v1.API
}

func NewAPI(urlService model.IURLService) *API {
	return &API{
		V1API: v1.NewAPI(urlService),
	}
}

func (a *API) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	a.V1API.RegisterRoutes(r)
	return r
}
