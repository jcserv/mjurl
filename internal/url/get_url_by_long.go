package url

import (
	"context"
	"errors"
	"strings"

	"github.com/jcserv/mjurl/model"
)

type GetURLByLong struct {
	long model.LongURL
}

var ErrLongURLEmpty = errors.New("long URL cannot be empty")

func NewGetURLByLong(long string) (*GetURLByLong, error) {
	c := &GetURLByLong{}
	s := strings.Trim(long, " ")
	if s == "" {
		return nil, ErrLongURLEmpty
	}

	c.long = model.LongURL(s)
	return c, nil
}

func (c *GetURLByLong) Execute(ctx context.Context, s model.IURLService) (*model.URL, error) {
	return s.GetURLByLong(ctx, c.long)
}
