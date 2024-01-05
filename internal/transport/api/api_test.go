package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jcserv/mjurl/model"
	"github.com/jcserv/mjurl/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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
		mockFunc  func(t *testing.T, s *mocks.MockIURLService)
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := mocks.NewMockIURLService(ctrl)

			if test.mockFunc != nil {
				test.mockFunc(t, s)
			}

			api := NewAPI(s)
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
		mockFunc  func(t *testing.T, s *mocks.MockIURLService)
		expected  APITestExpectations
	}{
		{
			name: "simple test",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodGet, APIV1URLPath+"/my-short", nil)
				return req
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().GetURLByShort(gomock.Any(), model.ShortURL("my-short")).Return(&model.URL{}, nil)
			},
			expected: APITestExpectations{
				code: http.StatusOK,
			},
		},
		{
			name: "short not provided, should return 404",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodGet, APIV1URLPath+"/", nil)
				return req
			},
			expected: APITestExpectations{
				code: http.StatusNotFound,
			},
		},
		{
			name: "service returns error, should return 500",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodGet, APIV1URLPath+"/my-short", nil)
				return req
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().GetURLByShort(gomock.Any(), model.ShortURL("my-short")).Return(nil, errors.New("uh oh spaghetti-o"))
			},
			expected: APITestExpectations{
				code: http.StatusInternalServerError,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := mocks.NewMockIURLService(ctrl)

			if test.mockFunc != nil {
				test.mockFunc(t, s)
			}

			api := NewAPI(s)
			r := api.RegisterRoutes()
			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.inputFunc(t))
			assert.Equal(t, test.expected.code, w.Code)
		})
	}
}
