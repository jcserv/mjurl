package url

import (
	"context"
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
			Long: model.LongURL(s),
		},
	}
	return command, nil
}

func (c *CreateShortURL) Execute(ctx context.Context, s model.IURLService) (*model.URL, error) {
	short_url, err := s.ShortenURL(ctx, c.insert_args.Long)
	if err != nil {
		return nil, err
	}
	c.insert_args.Short = short_url
	err = s.InsertURL(ctx, &c.insert_args)
	if err != nil {
		return nil, err
	}
	return &c.insert_args, nil
}
