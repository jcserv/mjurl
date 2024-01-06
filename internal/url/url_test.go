package url

import (
	"context"
	"testing"

	"github.com/jcserv/mjurl/model"
	"github.com/stretchr/testify/assert"
)

func Test_URLService_ShortenURL(t *testing.T) {
	tests := []struct {
		name     string
		input    model.LongURL
		expected error
	}{
		{
			name:     "simple test",
			input:    "",
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewURLService()
			actual := s.ShortenURL(context.Background(), test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_URLService_GetURLByShort(t *testing.T) {
	tests := []struct {
		name        string
		input       model.ShortURL
		expected    *model.URL
		expectedErr error
	}{
		{
			name:        "simple test",
			input:       "",
			expected:    nil,
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewURLService()
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
		expected    *model.URL
		expectedErr error
	}{
		{
			name:        "simple test",
			input:       "",
			expected:    nil,
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := NewURLService()
			actual, actualErr := s.GetURLByLong(context.Background(), test.input)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}
