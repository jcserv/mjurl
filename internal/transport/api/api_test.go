package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CreateURL(t *testing.T) {
	api := NewAPI()
	r := api.RegisterRoutes()
	req, _ := http.NewRequest(http.MethodPost, APIV1URLPath, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func Test_GetURL(t *testing.T) {
	api := NewAPI()
	r := api.RegisterRoutes()
	req, _ := http.NewRequest(http.MethodGet, APIV1URLPath, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}
