package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcserv/mjurl/internal/transport/api/httputil"
	"github.com/jcserv/mjurl/internal/url"
	"github.com/jcserv/mjurl/internal/utils/log"
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

func (a *API) RegisterRoutes(r *mux.Router) {
	r.HandleFunc(APIV1URLPath, a.ShortenURL()).Methods(http.MethodPost)
	r.HandleFunc(GetURLPath, a.GetURL()).Methods(http.MethodGet)
}

func (a *API) ShortenURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var input model.ShortenURLInput
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			httputil.BadRequest(w)
			return
		}

		command, err := url.NewShortenURL(input.URL)
		if err != nil {
			httputil.BadRequest(w)
			return
		}

		url, err := command.Execute(ctx, a.URLService)
		if err != nil {
			log.Error(ctx, fmt.Sprintf("unable to shorten url: %v", err))
			httputil.InternalServerError(ctx, w, err)
			return
		}
		httputil.OK(w, url.Short)
	}
}

func (a *API) GetURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)

		cmd, err := url.NewGetURLByShort(vars["short"])
		if err != nil {
			httputil.BadRequest(w)
			return
		}

		u, err := cmd.Execute(ctx, a.URLService)
		if err != nil {
			log.Error(ctx, fmt.Sprintf("unable to get url: %v", err))
			httputil.InternalServerError(ctx, w, err)
			return
		}
		httputil.PermanentRedirect(w, string(u.Long))
	}
}
