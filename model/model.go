package model

import (
	"bytes"
	"context"
	"encoding/json"
)

type URLID string
type ShortURL string
type LongURL string

type URL struct {
	ID    URLID    `json:"id"`
	Short ShortURL `json:"short_url"`
	Long  LongURL  `json:"long_url"`
}

func (u *URL) ToBytes() *bytes.Buffer {
	js, err := json.Marshal(u)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(js)
}

type IURLService interface {
	ShortenURL(ctx context.Context, long LongURL) error
	GetURLByShort(ctx context.Context, short ShortURL) (*URL, error)
	GetURLByLong(ctx context.Context, long LongURL) (*URL, error)
}

type IURLStore interface {
	CreateURL(ctx context.Context, url *URL) error
	ReadURL(ctx context.Context, params URLQueryParams) (*URL, error)
}

type URLQueryParams struct {
	URL
}
