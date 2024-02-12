package url

import (
	"context"
	"errors"
	"strings"

	"github.com/jcserv/mjurl/model"
)

type CreateShortURL struct {
	insert_args model.URL
}

func NewCreateShortURL(long string) (*CreateShortURL, error) {
	s := strings.Trim(long, " ")
	if s == "" {
		return nil, ErrLongURLEmpty
	}
	command := &CreateShortURL{
		insert_args: model.URL{
			Long: s
		}
	}
	return command, nil
}

func (c *CreateShortURL) Execute(ctx context.Context, s model.IURLService) error {
	c.insert_args.short_url, err := s.ShortenURL(ctx, c.insert_args.Long)
	if err != nil {
		return nil, err
	}
	return s.InsertURL(ctx, c)
}