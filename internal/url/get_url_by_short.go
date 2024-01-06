package url

import (
	"context"
	"errors"
	"strings"

	"github.com/jcserv/mjurl/model"
)

type GetURLByShort struct {
	short model.ShortURL
}

var ErrShortURLEmpty = errors.New("short URL cannot be empty")

func NewGetURLByShort(short string) (*GetURLByShort, error) {
	c := &GetURLByShort{}
	s := strings.Trim(short, " ")
	if s == "" {
		return nil, ErrShortURLEmpty
	}

	c.short = model.ShortURL(s)
	return c, nil
}

func (c *GetURLByShort) Execute(ctx context.Context, s model.IURLService) (*model.URL, error) {
	return s.GetURLByShort(ctx, c.short)
}
