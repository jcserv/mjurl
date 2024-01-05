// Code generated by MockGen. DO NOT EDIT.
// Source: internal/url/service.go
//
// Generated by this command:
//
//	mockgen -source=internal/url/service.go -destination=test/mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	url "github.com/jcserv/mjurl/internal/url"
	gomock "go.uber.org/mock/gomock"
)

// MockIURLService is a mock of IURLService interface.
type MockIURLService struct {
	ctrl     *gomock.Controller
	recorder *MockIURLServiceMockRecorder
}

// MockIURLServiceMockRecorder is the mock recorder for MockIURLService.
type MockIURLServiceMockRecorder struct {
	mock *MockIURLService
}

// NewMockIURLService creates a new mock instance.
func NewMockIURLService(ctrl *gomock.Controller) *MockIURLService {
	mock := &MockIURLService{ctrl: ctrl}
	mock.recorder = &MockIURLServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIURLService) EXPECT() *MockIURLServiceMockRecorder {
	return m.recorder
}

// GetURLByShort mocks base method.
func (m *MockIURLService) GetURLByShort(ctx context.Context, short url.ShortURL) (*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURLByShort", ctx, short)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURLByShort indicates an expected call of GetURLByShort.
func (mr *MockIURLServiceMockRecorder) GetURLByShort(ctx, short any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURLByShort", reflect.TypeOf((*MockIURLService)(nil).GetURLByShort), ctx, short)
}

// ShortenURL mocks base method.
func (m *MockIURLService) ShortenURL(ctx context.Context, long url.LongURL) (*url.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShortenURL", ctx, long)
	ret0, _ := ret[0].(*url.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShortenURL indicates an expected call of ShortenURL.
func (mr *MockIURLServiceMockRecorder) ShortenURL(ctx, long any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShortenURL", reflect.TypeOf((*MockIURLService)(nil).ShortenURL), ctx, long)
}