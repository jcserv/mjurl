package model

import (
	"context"
)

type URLID string
type ShortURL string
type LongURL string

type URL struct {
	ID    URLID    `json:"id"`
	Short ShortURL `json:"short_url"`
	Long  LongURL  `json:"long_url"`
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
