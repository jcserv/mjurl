package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/jcserv/mjurl/internal/transport/api/v1"
	"github.com/jcserv/mjurl/internal/utils/log"
	"github.com/jcserv/mjurl/model"
)

const (
	APIV1URLPath = "/api/v1/url"
	GetURLPath   = APIV1URLPath + "/{short}"
)

type API struct {
	V1API *v1.API
}

func NewAPI(dependencies model.Dependencies) *API {
	return &API{
		V1API: v1.NewAPI(dependencies.URLService),
	}
}

func LogIncomingRequests() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log.Info(ctx, fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		})
	}
}

func (a *API) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(LogIncomingRequests())
	a.V1API.RegisterRoutes(r)
	return r
}
