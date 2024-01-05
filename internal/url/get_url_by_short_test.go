package url

import (
	"context"
	"testing"

	"github.com/jcserv/mjurl/model"
	"github.com/jcserv/mjurl/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_NewGetURLByShort(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *GetURLByShort
		expectedErr error
	}{
		{
			name:  "happy path",
			input: "my-short-url",
			expected: &GetURLByShort{
				short: model.ShortURL("my-short-url"),
			},
			expectedErr: nil,
		},
		{
			name:        "input cannot be empty",
			input:       "",
			expected:    nil,
			expectedErr: ErrShortURLEmpty,
		},
		{
			name:        "input cannot be empty spaces",
			input:       "    ",
			expected:    nil,
			expectedErr: ErrShortURLEmpty,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, actualErr := NewGetURLBYShort(test.input)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}

func Test_GetURLByShort_Execute(t *testing.T) {
	tests := []struct {
		name        string
		input       *GetURLByShort
		mockFunc    func(t *testing.T, s *mocks.MockIURLService)
		expected    *model.URL
		expectedErr error
	}{
		{
			name: "simple test",
			input: &GetURLByShort{
				short: model.ShortURL("my-short-url"),
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().GetURLByShort(gomock.Any(), model.ShortURL("my-short-url")).Return(&model.URL{}, nil)
			},
			expected:    &model.URL{},
			expectedErr: nil,
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

			actual, actualErr := test.input.Execute(context.Background(), s)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		},
		)
	}
}
