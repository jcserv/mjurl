package url

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_URLService_GetURLByShort(t *testing.T) {
	tests := []struct {
		name        string
		input       ShortURL
		expected    *URL
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

func Test_URLService_ShortenURL(t *testing.T) {
	tests := []struct {
		name        string
		input       LongURL
		expected    *URL
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
			actual, actualErr := s.ShortenURL(context.Background(), test.input)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.expectedErr, actualErr)
		})
	}
}
