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
	ID    URLID    `json:"id" db:"id"`
	Short ShortURL `json:"short_url" db:"short_url"`
	Long  LongURL  `json:"long_url" db:"long_url"`
}

func (u *URL) ToBytes() *bytes.Buffer {
	js, err := json.Marshal(u)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(js)
}

type ShortenURLInput struct {
	URL string `json:"url"`
}

type IURLService interface {
	InsertURL(ctx context.Context, url *URL) error
	ShortenURL(ctx context.Context, long LongURL) (ShortURL, error)
	GetURLByShort(ctx context.Context, short ShortURL) (*URL, error)
	GetURLByLong(ctx context.Context, long LongURL) (*URL, error)
}

type IURLStore interface {
	CreateURL(ctx context.Context, url *URL) error
	QueryURL(ctx context.Context, id URLID) (*URL, error)
	QueryURLByShort(ctx context.Context, short ShortURL) (*URL, error)
	QueryURLByLong(ctx context.Context, long LongURL) (*URL, error)
}

type Dependencies struct {
	URLService IURLService
}
