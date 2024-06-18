package url

import (
	"context"
	"testing"

	"github.com/jcserv/mjurl/model"
	"github.com/jcserv/mjurl/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_URLService_ShortenURL(t *testing.T) {
	tests := []struct {
		name        string
		input       model.LongURL
		expected    model.ShortURL
		expectedErr error
	}{
		{
			name:     "simple test",
			input:    "",
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			urlStore := mocks.NewMockIURLStore(ctrl)
			s := NewURLService(urlStore)
			actual, actualErr := s.ShortenURL(context.Background(), test.input)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}

func Test_URLService_GetURLByShort(t *testing.T) {
	tests := []struct {
		name        string
		input       model.ShortURL
		mockFunc    func(t *testing.T, s *mocks.MockIURLStore)
		expected    *model.URL
		expectedErr error
	}{
		{
			name:  "happy path",
			input: mocks.URL.Short,
			mockFunc: func(t *testing.T, s *mocks.MockIURLStore) {
				t.Helper()
				s.EXPECT().QueryURLByShort(gomock.Any(), mocks.URL.Short).Return(mocks.URL, nil)
			},
			expected:    mocks.URL,
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			urlStore := mocks.NewMockIURLStore(ctrl)

			if test.mockFunc != nil {
				test.mockFunc(t, urlStore)
			}

			s := NewURLService(urlStore)
			actual, actualErr := s.GetURLByShort(context.Background(), test.input)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}

func Test_URLService_GetURLByLong(t *testing.T) {
	tests := []struct {
		name        string
		input       model.LongURL
		mockFunc    func(t *testing.T, s *mocks.MockIURLStore)
		expected    *model.URL
		expectedErr error
	}{
		{
			name:  "happy path",
			input: mocks.URL.Long,
			mockFunc: func(t *testing.T, s *mocks.MockIURLStore) {
				t.Helper()
				s.EXPECT().QueryURLByLong(gomock.Any(), mocks.URL.Long).Return(mocks.URL, nil)
			},
			expected:    mocks.URL,
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			urlStore := mocks.NewMockIURLStore(ctrl)

			if test.mockFunc != nil {
				test.mockFunc(t, urlStore)
			}

			s := NewURLService(urlStore)
			actual, actualErr := s.GetURLByLong(context.Background(), test.input)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}
