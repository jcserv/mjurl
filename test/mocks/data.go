package mocks

import (
	"errors"

	"github.com/jcserv/mjurl/model"
)

const (
	ShortURL = "my-short-url"
	LongURL  = "my-long-url"
)

var URL = &model.URL{
	ID:    "2009215674938",
	Short: "zn9edcu",
	Long:  "https://en.wikipedia.org/wiki/Systems_design",
}

var Err = errors.New("mock err")
