package url

import "context"

type URLID string
type ShortURL string
type LongURL string

type URL struct {
	ID    URLID    `json:"id"`
	Short ShortURL `json:"short_url"`
	Long  LongURL  `json:"long_url"`
}

type IURLService interface {
	GetURLByShort(ctx context.Context, short ShortURL) (*URL, error)
	ShortenURL(ctx context.Context, long LongURL) (*URL, error)
}

type URLService struct {
}

func NewURLService() IURLService {
	return &URLService{}
}

// GetURLByShort gets the URL object associated with the given short URL.
func (s *URLService) GetURLByShort(ctx context.Context, short ShortURL) (*URL, error) {
	return nil, nil
}

// ShortenURL is a function that generates a URL object for a given URL.
func (s *URLService) ShortenURL(ctx context.Context, long LongURL) (*URL, error) {
	return nil, nil
}
