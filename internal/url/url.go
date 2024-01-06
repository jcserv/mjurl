package url

import (
	"context"

	"github.com/jcserv/mjurl/model"
)

type URLService struct {
}

func NewURLService() model.IURLService {
	return &URLService{}
}

// ShortenURL is a function that generates a URL object for a given URL.
func (s *URLService) ShortenURL(ctx context.Context, long model.LongURL) error {
	return nil
}

// GetURLByShort gets the URL object associated with the given short URL.
func (s *URLService) GetURLByShort(ctx context.Context, short model.ShortURL) (*model.URL, error) {
	return nil, nil
}

// GetURLByLong gets the URL object associated with the given long URL.
func (s *URLService) GetURLByLong(ctx context.Context, long model.LongURL) (*model.URL, error) {
	return nil, nil
}
