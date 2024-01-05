package url

import (
	"context"

	"github.com/jcserv/mjurl/model"
)

type IURLService interface {
	GetURLByShort(ctx context.Context, short model.ShortURL) (*model.URL, error)
	ShortenURL(ctx context.Context, long model.LongURL) (*model.URL, error)
}

type URLService struct {
}

func NewURLService() IURLService {
	return &URLService{}
}

// GetURLByShort gets the URL object associated with the given short URL.
func (s *URLService) GetURLByShort(ctx context.Context, short model.ShortURL) (*model.URL, error) {
	return nil, nil
}

// ShortenURL is a function that generates a URL object for a given URL.
func (s *URLService) ShortenURL(ctx context.Context, long model.LongURL) (*model.URL, error) {
	return nil, nil
}
