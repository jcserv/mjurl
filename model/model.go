package model

type URLID string
type ShortURL string
type LongURL string

type URL struct {
	ID    URLID    `json:"id"`
	Short ShortURL `json:"short_url"`
	Long  LongURL  `json:"long_url"`
}
