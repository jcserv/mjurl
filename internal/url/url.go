package url

import (
	"context"
	"crypto/md5"

	"github.com/jcserv/mjurl/model"
)

type URLService struct {
	urlStore model.IURLStore
}

func NewURLService(urlStore model.IURLStore) model.IURLService {
	return &URLService{urlStore}
}

// ShortenURL is a function that generates a URL object for a given URL.
func (s *URLService) ShortenURL(ctx context.Context, long model.LongURL) (model.ShortURL, error) {
	h := md5.New()
	h.Write([]byte(long))
	hashValue := h.Sum(nil)
	return model.ShortURL(hashValue), nil
}

func (s *URLService) InsertURL(ctx context.Context, url *model.URL) error {
	return s.urlStore.CreateURL(ctx, url)
}

// GetURLByShort gets the URL object associated with the given short URL.
func (s *URLService) GetURLByShort(ctx context.Context, short model.ShortURL) (*model.URL, error) {
	return s.urlStore.QueryURLByShort(ctx, short)
}

// GetURLByLong gets the URL object associated with the given long URL.
func (s *URLService) GetURLByLong(ctx context.Context, long model.LongURL) (*model.URL, error) {
	return s.urlStore.QueryURLByLong(ctx, long)
}
