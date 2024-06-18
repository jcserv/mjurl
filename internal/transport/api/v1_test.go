package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jcserv/mjurl/model"
	"github.com/jcserv/mjurl/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type V1APITest struct {
	name       string
	inputFunc  func(t *testing.T) *http.Request
	mockFunc   func(t *testing.T, s *mocks.MockIURLService)
	assertFunc func(t *testing.T, expected, actual APITestExpectations)
	expected   APITestExpectations
}

type APITestExpectations struct {
	code         int
	header       http.Header
	responseBody any
}

func Test_ShortenURL(t *testing.T) {
	expectedResp, _ := json.Marshal(mocks.URL.Short)
	tests := []V1APITest{
		{
			name: "simple test",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				input := model.ShortenURLInput{
					URL: string(mocks.URL.Long),
				}
				jsonData, _ := json.Marshal(input)
				req, _ := http.NewRequest(http.MethodPost, APIV1URLPath, bytes.NewReader(jsonData))
				return req
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().ShortenURL(gomock.Any(), model.LongURL(mocks.URL.Long)).
					Return(mocks.URL.Short, nil)
				s.EXPECT().InsertURL(gomock.Any(), mocks.URLWithoutID).
					Return(nil)
			},
			assertFunc: assertResponse,
			expected: APITestExpectations{
				code: http.StatusOK,
				header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				responseBody: bytes.NewBuffer(expectedResp),
			},
		},
		{
			name: "invalid request body, should return 400",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				input := model.ShortenURLInput{
					URL: "",
				}
				jsonData, _ := json.Marshal(input)
				req, _ := http.NewRequest(http.MethodPost, APIV1URLPath, bytes.NewReader(jsonData))
				return req
			},
			assertFunc: assertStatusCode,
			expected: APITestExpectations{
				code: http.StatusBadRequest,
			},
		},
		{
			name: "service function returns error, should return 500",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				input := model.ShortenURLInput{
					URL: string(mocks.URL.Long),
				}
				jsonData, _ := json.Marshal(input)
				req, _ := http.NewRequest(http.MethodPost, APIV1URLPath, bytes.NewReader(jsonData))
				return req
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().ShortenURL(gomock.Any(), model.LongURL(mocks.URL.Long)).
					Return(model.ShortURL(""), mocks.Err)
			},
			assertFunc: assertStatusCode,
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

			api := NewAPI(Dependencies{
				URLService: s,
			})
			r := api.RegisterRoutes()
			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.inputFunc(t))

			result := w.Result()
			actual := APITestExpectations{
				code:         result.StatusCode,
				header:       w.Header(),
				responseBody: w.Body,
			}

			test.assertFunc(t, test.expected, actual)
		})
	}
}

func Test_GetURL(t *testing.T) {
	tests := []V1APITest{
		{
			name: "simple test",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", APIV1URLPath, mocks.URL.Short), nil)
				return req
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().GetURLByShort(gomock.Any(), model.ShortURL(mocks.URL.Short)).
					Return(mocks.URL, nil)
			},
			assertFunc: assertStatusAndLocationHeader,
			expected: APITestExpectations{
				code: http.StatusPermanentRedirect,
				header: http.Header{
					"Content-Type": []string{"application/json"},
					"Location":     []string{string(mocks.URL.Long)},
				},
			},
		},
		{
			name: "short not provided, should return 404",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodGet, APIV1URLPath+"/", nil)
				return req
			},
			assertFunc: assertStatusCode,
			expected: APITestExpectations{
				code: http.StatusNotFound,
			},
		},
		{
			name: "service returns error, should return 500",
			inputFunc: func(t *testing.T) *http.Request {
				t.Helper()
				req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", APIV1URLPath, mocks.ShortURL), nil)
				return req
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().GetURLByShort(gomock.Any(), model.ShortURL(mocks.ShortURL)).
					Return(nil, mocks.Err)
			},
			assertFunc: assertStatusCode,
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

			api := NewAPI(Dependencies{
				URLService: s,
			})
			r := api.RegisterRoutes()
			w := httptest.NewRecorder()
			r.ServeHTTP(w, test.inputFunc(t))

			result := w.Result()
			actual := APITestExpectations{
				code:         result.StatusCode,
				header:       w.Header(),
				responseBody: w.Body,
			}

			test.assertFunc(t, test.expected, actual)
		})
	}
}

func assertResponse(t *testing.T, expected, actual APITestExpectations) {
	t.Helper()
	assert.Equal(t, expected, actual)
}

func assertStatusCode(t *testing.T, expected, actual APITestExpectations) {
	t.Helper()
	assert.Equal(t, expected.code, actual.code)
}

func assertStatusAndLocationHeader(t *testing.T, expected, actual APITestExpectations) {
	t.Helper()
	assert.Equal(t, expected.code, actual.code)
	assert.Equal(t, expected.header.Get("Location"), actual.header.Get("Location"))
}
