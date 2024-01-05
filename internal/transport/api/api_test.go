package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type APITestExpectations struct {
	code         int
	header       http.Header
	responseBody any
}

func Test_CreateURL(t *testing.T) {
	tests := []struct {
		name      string
		inputFunc func(t *testing.T) *http.Request
		expected  APITestExpectations
	}{
		{
			name: "simple test",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodPost, APIV1URLPath, nil)
				return req
			},
			expected: APITestExpectations{
				code: http.StatusOK,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			api := NewAPI()
			r := api.RegisterRoutes()
			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.inputFunc(t))
			assert.Equal(t, test.expected.code, w.Code)
		})
	}
}

func Test_GetURL(t *testing.T) {
	tests := []struct {
		name      string
		inputFunc func(t *testing.T) *http.Request
		expected  APITestExpectations
	}{
		{
			name: "simple test",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodGet, APIV1URLPath, nil)
				return req
			},
			expected: APITestExpectations{
				code: http.StatusOK,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			api := NewAPI()
			r := api.RegisterRoutes()
			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.inputFunc(t))
			assert.Equal(t, test.expected.code, w.Code)
		})
	}
}
