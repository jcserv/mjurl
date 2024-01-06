package url

import (
	"context"
	"testing"

	"github.com/jcserv/mjurl/model"
	"github.com/jcserv/mjurl/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_NewGetURLByLong(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *GetURLByLong
		expectedErr error
	}{
		{
			name:  "happy path",
			input: mocks.LongURL,
			expected: &GetURLByLong{
				long: model.LongURL(mocks.LongURL),
			},
			expectedErr: nil,
		},
		{
			name:        "input cannot be empty",
			input:       "",
			expected:    nil,
			expectedErr: ErrLongURLEmpty,
		},
		{
			name:        "input cannot be empty spaces",
			input:       "    ",
			expected:    nil,
			expectedErr: ErrLongURLEmpty,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, actualErr := NewGetURLByLong(test.input)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}

func Test_GetURLByLong_Execute(t *testing.T) {
	tests := []struct {
		name        string
		input       *GetURLByLong
		mockFunc    func(t *testing.T, s *mocks.MockIURLService)
		expected    *model.URL
		expectedErr error
	}{
		{
			name: "simple test",
			input: &GetURLByLong{
				long: model.LongURL(mocks.LongURL),
			},
			mockFunc: func(t *testing.T, s *mocks.MockIURLService) {
				t.Helper()
				s.EXPECT().GetURLByLong(gomock.Any(), model.LongURL(mocks.LongURL)).Return(&model.URL{}, nil)
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
