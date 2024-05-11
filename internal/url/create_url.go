package url

import (
	"context"
	"strings"

	"github.com/jcserv/mjurl/model"
)

type ShortenURL struct {
	url *model.URL
}

func NewShortenURL(long string) (*ShortenURL, error) {
	s := strings.Trim(long, " ")
	if s == "" {
		return nil, ErrLongURLEmpty
	}
	command := &ShortenURL{
		url: &model.URL{
			Long: model.LongURL(s),
		},
	}
	return command, nil
}

func (c *ShortenURL) Execute(ctx context.Context, s model.IURLService) (*model.URL, error) {
	shortURL, err := s.ShortenURL(ctx, c.url.Long)
	if err != nil {
		return nil, err
	}
	c.url.Short = shortURL

	err = s.InsertURL(ctx, c.url)
	if err != nil {
		return nil, err
	}

	return c.url, nil
}
